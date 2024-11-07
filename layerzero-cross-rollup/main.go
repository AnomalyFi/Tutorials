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
	"time"

	"github.com/bianyuanop/lz-cross-rollup/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/flashbotsrpc"
)

var (
	javelinUrl    = flag.String("javelin", getEnvAsStrOrDefault("JAVELIN_URL", "http://127.0.0.1:3000/rpc"), "rpc url of javelin rpc")
	gethUrlOrigin = flag.String("geth-origin", getEnvAsStrOrDefault("GETH_URL_ORIGIN", "http://127.0.0.1:9090"), "geth rpc url on origin chain")
	gethUrlRemote = flag.String("geth-remote", getEnvAsStrOrDefault("GETH_URL_REMOTE", "http://127.0.0.1:9091"), "geth rpc url on remote chain")
	originChainID = flag.Int("origin-chainid", getEnvAsIntOrDefault("ORIGIN_CHAINID", 45205), "chain id of origin chain")
	remoteChainID = flag.Int("remote-chainid", getEnvAsIntOrDefault("REMOTE_CHAINID", 45206), "chain id of remote chain")
	oftOrigin     = flag.String("oft-origin", getEnvAsStrOrDefault("OFT_ORIGIN", "0x985060F8b809F08392FB4E23622E9E6881c22d0b"), "oft address on origin chain")
	oftRemote     = flag.String("oft-remote", getEnvAsStrOrDefault("OFT_REMOTE", "0x985060F8b809F08392FB4E23622E9E6881c22d0b"), "oft address on remote chain")
	amount        = flag.Int("amount", getEnvAsIntOrDefault("AMOUNT", 1e+18), "amount of oft to transfer")
	nonceInc      = flag.Int("nonce-inc", getEnvAsIntOrDefault("NONCE_INC", 0), "the nonce adjustment applied to the latest nonce") // this is for testing multiple hyperlane txs from the same account

	privKey = flag.String("priv-key", getEnvAsStrOrDefault("PRIV_KEY", "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"), "priv key of wallet without 0x prefix")
)

func main() {
	flag.Parse()

	log.Printf("chainID: %d, %d\n", *originChainID, *remoteChainID)

	originClient, err := ethclient.Dial(*gethUrlOrigin)
	if err != nil {
		log.Fatal(err)
	}
	// remoteClient, err := ethclient.Dial(*gethUrlRemote)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var oftOriginAddress, oftRemoteAddress common.Address
	err = oftOriginAddress.UnmarshalText([]byte(*oftOrigin))
	if err != nil {
		log.Fatal(err)
	}

	err = oftRemoteAddress.UnmarshalText([]byte(*oftRemote))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("oft origin: %+v\n", oftOriginAddress)
	fmt.Printf("oft remote: %+v\n", oftRemoteAddress)

	privKey, err := crypto.HexToECDSA(*privKey)
	if err != nil {
		log.Fatal(err)
	}
	pubkey := privKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*pubkeyECDSA)
	nonce, err := originClient.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalf("unable to fetch nonce: %+v\n", err)
	}
	nonce += uint64(*nonceInc)

	gasPrice, err := originClient.SuggestGasPrice(context.TODO())
	if err != nil {
		log.Fatalf("unable to fetch gas price: %+v\n", err)
	}

	originChainID := big.NewInt(int64(*originChainID))
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, originChainID)
	if err != nil {
		log.Fatalf("unable to gen auth: %+v\n", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)         // on origin chain this has to be the exact amount you want to transfer to the token router contract
	auth.GasLimit = uint64(21000 * 10) // in units
	auth.GasPrice = gasPrice
	auth.NoSend = true

	oftOrigin, err := contracts.NewMyOFT(oftOriginAddress, originClient)
	if err != nil {
		log.Fatal(fmt.Errorf("contract err: %w", err))
	}
	extraOptions := contracts.NewOptions()
	extraOptions, err = contracts.AddExecutorLzReceiveOption(extraOptions, 65000, 0)
	if err != nil {
		log.Fatal(err)
	}
	sendParams := contracts.SendParam{
		DstEid:       uint32(*remoteChainID),
		To:           padAddressToBytes32(oftRemoteAddress),
		AmountLD:     big.NewInt(int64(*amount)),
		MinAmountLD:  big.NewInt(int64(*amount * 9 / 10)),
		ExtraOptions: extraOptions,
		ComposeMsg:   nil,
		OftCmd:       nil,
	}
	fee, err := oftOrigin.QuoteSend(nil, sendParams, false)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to quote: %w", err))
	}
	fmt.Printf("quote fee: %+v\n", fee)
	tx, err := oftOrigin.Send(auth, sendParams, fee, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	txRaw, err := tx.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	txs := make(map[string]string)
	txs[hexutil.EncodeBig(originChainID)] = hexutil.Encode(txRaw)

	blockNumber := make(map[string]string)
	blockNumber[hexutil.EncodeBig(originChainID)] = "0x0"

	bundleArgs := flashbotsrpc.FlashbotsSendBundleCrossRollupRequest{
		Txs:         txs,
		BlockNumber: blockNumber,
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

	fmt.Printf("tx receipt: %+v\n", txReceipt)
}

func padAddressToBytes32(addr common.Address) [32]byte {
	var ret [32]byte
	copy(ret[12:], addr[:])

	return ret
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
			switch bundleStatusResp.StatusCode {
			case 0x0:
				log.Printf("bundle accepted")
				return nil
			case 0x1:
				log.Fatalf("bundle rejected, %+v", bundleStatusResp.Status)
			case 0x2:
				log.Printf("bundle not sent to SEQ yet, Status: %s", bundleStatusResp.Status)
			}
		}
	}
}

func balanceOfNative(ethcli *ethclient.Client, account common.Address) (*big.Int, error) {
	return ethcli.BalanceAt(context.Background(), account, nil)
}

func balanceOfERC20(ethcli *ethclient.Client, contractAddress common.Address, account common.Address) (*big.Int, error) {
	erc20, err := contracts.NewMyOFT(contractAddress, ethcli)
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