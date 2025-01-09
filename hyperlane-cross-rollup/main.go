package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bianyuanop/hyperlane-demo/contracts/hyperc20"
	"github.com/bianyuanop/hyperlane-demo/contracts/mailbox"
	"github.com/bianyuanop/hyperlane-demo/contracts/tokenrouter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/flashbotsrpc"
)

// In javelin-rpc, we have two paths for atomic inclusion
// 1. cross rollup bundle, which is introduced in the other tutorial
// 2. hyperlane path, which can bridge tokens back and forth, which relies on Hyperlane
// for detail, visit https://docs.hyperlane.xyz/docs/protocol/warp-routes
// unlike what's on rollups based on hyperlane, where you have to issue the transferRemote tx first on origin chain
// after some time, the validators verify and create a multi-sig on it, the relayer will populate a remote tx to be sent on remote chain to achieve token transfer
// here in Nodekit-SEQ based rollups, we verify one hyperlane tx validity and populate the remote tx right away to simulate then send them as a cross-rollup bundle
// to achieve atomic inclusion

// In the devnet at Jun-26, we deployed HypNative contract on origin chain(45206) and HypERC20 on remote chain(45207) to issue transfers
// there are two directions transfer
// native -> synthetic (45206 -> 45207) or
// synthetic -> native (45207 -> 45206)

var (
	javelinUrl        = flag.String("javelin", getEnvAsStrOrDefault("JAVELIN_URL", "http://127.0.0.1:3000/rpc"), "rpc url of javelin rpc")
	gethUrlOrigin     = flag.String("geth-origin", getEnvAsStrOrDefault("GETH_URL_ORIGIN", "http://127.0.0.1:19551"), "geth rpc url on origin chain")
	gethUrlRemote     = flag.String("geth-remote", getEnvAsStrOrDefault("GETH_URL_REMOTE", "http://127.0.0.1:19552"), "geth rpc url on remote chain")
	originChainID     = flag.Int("origin-chainid", getEnvAsIntOrDefault("ORIGIN_CHAINID", 45206), "chain id of origin chain")
	remoteChainID     = flag.Int("remote-chainid", getEnvAsIntOrDefault("REMOTE_CHAINID", 45207), "chain id of remote chain")
	mailboxOrigin     = flag.String("mailbox-origin", getEnvAsStrOrDefault("MAILBOX_ORIGIN", "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"), "mailbox address on origin chain")
	mailboxRemote     = flag.String("mailbox-remote", getEnvAsStrOrDefault("MAILBOX_REMOTE", "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"), "mailbox address on remote chain")
	tokenType         = flag.String("token-type", getEnvAsStrOrDefault("TOKEN_TYPE", "native"), "token type on origin chain to transfer from origin chain to remote chain")
	tokenRouterOrigin = flag.String("router-origin", getEnvAsStrOrDefault("TOKEN_ROUTER_ORIGIN", "0x4A679253410272dd5232B3Ff7cF5dbB88f295319"), "token router address on origin chain")
	tokenRouterRemote = flag.String("router-remote", getEnvAsStrOrDefault("TOKEN_ROUTER_REMOTE", "0x4A679253410272dd5232B3Ff7cF5dbB88f295319"), "token router address on remote chain")
	target            = flag.String("target", getEnvAsStrOrDefault("TARGET", "0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"), "which account to send token to")
	amount            = flag.Int("amount", getEnvAsIntOrDefault("AMOUNT", 1), "amount of token to transfer")
	nonce             = flag.Int("nonce", getEnvAsIntOrDefault("NONCE_INC", 0), "the nonce adjustment applied to the latest nonce") // this is for testing multiple hyperlane txs from the same account
	timestamp         = flag.Int("timestamp", getEnvAsIntOrDefault("TIMESTAMP", 0), "timestamp for ordering of bundles on javelin-rpc")

	privKey = flag.String("priv-key", getEnvAsStrOrDefault("PRIV_KEY", "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"), "priv key of wallet without 0x prefix")
)

func main() {
	flag.Parse()

	log.Printf("chainID: %d, %d\n", *originChainID, *remoteChainID)

	originClient, err := ethclient.Dial(*gethUrlOrigin)
	if err != nil {
		log.Fatal(err)
	}
	remoteClient, err := ethclient.Dial(*gethUrlRemote)
	if err != nil {
		log.Fatal(err)
	}

	// ======== populate transferRemote tx on origin chain =============
	tokenRouterOriginAddress := new(common.Address)
	err = tokenRouterOriginAddress.UnmarshalText([]byte(*tokenRouterOrigin))
	if err != nil {
		log.Fatalf("unable to unmarshal router origin addr: %+v\n", err)
	}
	tokenRouterOrigin, err := tokenrouter.NewTokenRouter(*tokenRouterOriginAddress, originClient)
	if err != nil {
		log.Fatalf("unable to new TokenRouter at origin: %+v\n", err)
	}

	privKey, err := crypto.HexToECDSA(strings.TrimLeft(*privKey, "0x"))
	if err != nil {
		log.Fatal(err)
	}
	pubkey := privKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*pubkeyECDSA)
	nonce2use := uint64(*nonce)
	if *nonce == 0 {
		nonceQueried, err := originClient.NonceAt(context.Background(), fromAddress, nil)
		if err != nil {
			log.Fatalf("unable to fetch nonce: %+v\n", err)
		}
		nonce2use = nonceQueried
	}

	gasPrice, err := originClient.SuggestGasPrice(context.TODO())
	if err != nil {
		log.Fatalf("unable to fetch gas price: %+v\n", err)
	}

	// interchain gas payment + amount to transfer
	gasPayment, err := tokenRouterOrigin.QuoteGasPayment(nil, uint32(*remoteChainID))
	if err != nil {
		log.Fatalf("unable to quote gas payment: %+v\n", err)
	}

	txValue := gasPayment
	amount2transfer := big.NewInt(int64(*amount))
	if *tokenType == "native" {
		// is native balance sufficient to cover IGP(interchain gas payment) and native token transfer
		txValue = txValue.Add(txValue, amount2transfer)
		bal, err := balanceOfNative(originClient, fromAddress)
		if err != nil {
			log.Fatalf("unable to query balance of account: %+v\n", err)
		}
		// insufficient balance
		if bal.Cmp(txValue) < 0 {
			log.Fatalf("insufficient balance to cover transfer, want: %+v, actual: %+v\n", txValue, bal)
		}
	} else if *tokenType == "synthetic" || *tokenType == "erc20" {
		// to check if balance of erc20 is sufficient & native token balance can cover inter chain gas
		erc20bal, err := balanceOfERC20(originClient, *tokenRouterOriginAddress, fromAddress)
		if err != nil {
			log.Fatalf("unable to get erc20 balance of account: %+v\n", err)
		}
		if erc20bal.Cmp(amount2transfer) < 0 {
			log.Fatalf("insufficient balance to cover erc20/synthetic transfer, want: %+v, actual: %+v\n", amount2transfer, erc20bal)
		}

		bal, err := balanceOfNative(originClient, fromAddress)
		if err != nil {
			log.Fatalf("unable to query balance of account: %+v\n", err)
		}
		// insufficient balance
		if bal.Cmp(txValue) < 0 {
			log.Fatalf("insufficient balance to pay gas, want: %+v, actual: %+v\n", txValue, bal)
		}
	} else {
		log.Fatalf("unsupported token type: %s\n", *tokenType)
	}

	originChainID := big.NewInt(int64(*originChainID))
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, originChainID)
	if err != nil {
		log.Fatalf("unable to gen auth: %+v\n", err)
	}
	auth.Nonce = big.NewInt(int64(nonce2use))
	auth.Value = txValue           // on origin chain this has to be the exact amount you want to transfer to the token router contract
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.NoSend = true

	targetAddress := common.FromHex(*target)
	// pad 0 to fill up the 32 bytes
	var remoteAddress32 [32]byte
	remoteAddress := hexutil.MustDecode(string(targetAddress))
	copy(remoteAddress32[32-len(remoteAddress):], remoteAddress)
	log.Printf("address32: %s, address: %s\n", hexutil.Encode(remoteAddress32[:]), hexutil.Encode(remoteAddress))

	tx, err := tokenRouterOrigin.TransferRemote(auth, uint32(*remoteChainID), remoteAddress32, amount2transfer)
	if err != nil {
		log.Fatalf("unable to populate transfer remote: %+v\n", err)
	}

	log.Printf("transfer remote tx has been populated, nonce: %d, sender: %s, chainID: %d\n", tx.Nonce(), fromAddress.Hex(), tx.ChainId().Int64())

	// construct javelin cross-rollup bundle to send, we allow the bundle only contain 1 tx
	txRaw, err := tx.MarshalBinary()
	if err != nil {
		log.Fatalf("unable to marshal tx to binary: %s", err)
	}
	txs := make(map[string]string)
	txs[hexutil.EncodeBig(originChainID)] = hexutil.Encode(txRaw)

	blockNumber := make(map[string]string)
	blockNumber[hexutil.EncodeBig(originChainID)] = "0x0"

	minTimestamp := (*uint64)(nil)
	if *timestamp != 0 {
		ts := uint64(*timestamp)
		minTimestamp = &ts
	}
	bundleArgs := flashbotsrpc.FlashbotsSendBundleCrossRollupRequest{
		Txs:          txs,
		BlockNumber:  blockNumber,
		MinTimestamp: minTimestamp,
	}

	fbRpc := flashbotsrpc.NewFlashbotsRPC(*javelinUrl)
	resp, err := fbRpc.FlashbotsSendBundleCrossRollup(privKey, bundleArgs)
	if err != nil {
		log.Fatalf("unable to send cross-rollup bundle: %s", err)
	}
	bundleHash := resp.BundleHash
	log.Printf("submitted bundle: %s\n", bundleHash)

	ctx := context.Background()
	bctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	if err := waitForBundleAcceptance(bctx, privKey, fbRpc, bundleHash); err != nil {
		log.Fatalf("bundle not accepted: %s", err)
	}
	cancel()

	tctx, cancel := context.WithTimeout(ctx, 180*time.Second)
	defer cancel()
	txReceipt, err := waitForReceipt(tctx, originClient, tx.Hash())
	if err != nil {
		log.Fatalf("unable to get tx receipt: %+v\n", err)
	}

	mailboxAddressOrigin := new(common.Address)
	err = mailboxAddressOrigin.UnmarshalText([]byte(*mailboxOrigin))
	if err != nil {
		log.Fatalf("unable to unmarshal mailbox addr on origin: %+v\n", err)
	}
	mailboxOrigin, err := mailbox.NewMailbox(*mailboxAddressOrigin, originClient)
	if err != nil {
		log.Fatalf("unable to create mailbox instance: %+v\n", err)
	}

	// get dispatched msg ids to query on remote chain
	messageIDs := make([]common.Hash, 0)
	for _, txLog := range txReceipt.Logs {
		dispatchedMsg, err := mailboxOrigin.ParseDispatch(*txLog)
		if err != nil {
			// log.Printf("unable to parse dispatch message for log %+v, err(%+v)\n", txLog, err)
			continue
		}
		messageID := crypto.Keccak256Hash(dispatchedMsg.Message)
		messageIDs = append(messageIDs, messageID)

		log.Printf("Dipsatch message(%+v) detected: %+v\n", messageID, txLog)
	}

	log.Printf("num of dispatchMsg detected: %d\n", len(messageIDs))

	mailboxAddressRemote := new(common.Address)
	err = mailboxAddressRemote.UnmarshalText([]byte(*mailboxRemote))
	if err != nil {
		log.Fatalf("unable to unmarshal mailbox addr on remote: %+v\n", err)
	}
	mailboxRemote, err := mailbox.NewMailbox(*mailboxAddressRemote, remoteClient)
	if err != nil {
		log.Fatalf("unable to create mailbox instance: %+v\n", err)
	}
	for _, msgID := range messageIDs {
		mctx, cancel := context.WithTimeout(ctx, 180*time.Second)
		defer cancel()
		err := waitMsgToBeProcessed(mctx, msgID, mailboxRemote)
		if err != nil {
			log.Printf("timeout waiting msg(%+v) to be processed", msgID)
		}
		log.Printf("message processed: %+v\n", msgID)
	}
	log.Printf("funds has beed moved from chain %+v to chain %+v\n", *originChainID, *remoteChainID)
}

func waitMsgToBeProcessed(ctx context.Context, msgID common.Hash, mailbox *mailbox.Mailbox) error {
	log.Printf("waiting for message(%+v) to be processed\n", msgID)
	for {
		select {
		case <-time.After(1 * time.Second):
			_, err := mailbox.Delivered(nil, [32]byte(msgID.Bytes()))
			if err != nil {
				log.Printf("unable to query delivery status: %+v\n", err)
				continue
			}
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func waitForReceipt(ctx context.Context, ethcli *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	log.Printf("waiting for receipt of tx: %s\n", txHash.Hex())
	for {
		select {
		case <-time.After(1 * time.Second):
			receipt, err := ethcli.TransactionReceipt(ctx, txHash)
			if err != nil {
				continue
			}
			return receipt, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func waitForBundleAcceptance(ctx context.Context, pk *ecdsa.PrivateKey, rpc *flashbotsrpc.FlashbotsRPC, bundleHash string) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("bundle not accepted: %s", ctx.Err())
		default:
			time.Sleep(500 * time.Millisecond)
			bundleStatusResp, err := rpc.FlashbotsGetBundleCrossRollupStatus(pk, flashbotsrpc.FlashbotsGetBundleCrossRollupStatusRequest{
				BundleHash: bundleHash,
			})
			if err != nil {
				log.Printf("error issuing eth_getBundleCrossRollup: %+v\n", err)
				continue
			}
			log.Printf("status resp: %+v\n", bundleStatusResp)
			// included
			switch bundleStatusResp.Status {
			case "accepted":
				log.Printf("bundle accepted")
				return nil
			case "rejected":
				log.Fatalf("bundle rejected, %s", bundleStatusResp.Status)
			case "included":
				log.Fatalf("bundle included, %s", bundleStatusResp.Status)
			default:
				log.Printf("unknown bundle status: %s", bundleStatusResp.Status)
			}
		}
	}
}

func balanceOfNative(ethcli *ethclient.Client, account common.Address) (*big.Int, error) {
	return ethcli.BalanceAt(context.Background(), account, nil)
}

func balanceOfERC20(ethcli *ethclient.Client, contractAddress common.Address, account common.Address) (*big.Int, error) {
	erc20, err := hyperc20.NewHypERC20(contractAddress, ethcli)
	if err != nil {
		return nil, err
	}

	return erc20.BalanceOf(nil, account)
}

func getEnvAsStrOrDefault(key string, defaultValue string) string {
	ret := os.Getenv(key)
	if ret == "" {
		ret = defaultValue
	}
	return ret
}

func getEnvAsIntOrDefault(name string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
