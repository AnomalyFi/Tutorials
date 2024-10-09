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
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var (
	javelinUrl    = flag.String("javelin", getEnvAsStrOrDefault("JAVELIN_URL", "http://127.0.0.1:3000/rpc"), "rpc url of javelin rpc")
	gethUrlOrigin = flag.String("geth-origin", getEnvAsStrOrDefault("GETH_URL_ORIGIN", "http://127.0.0.1:19550"), "geth rpc url on origin chain")
	gethUrlRemote = flag.String("geth-remote", getEnvAsStrOrDefault("GETH_URL_REMOTE", "http://127.0.0.1:19551"), "geth rpc url on remote chain")
	originChainID = flag.Int("origin-chainid", getEnvAsIntOrDefault("ORIGIN_CHAINID", 45205), "chain id of origin chain")
	remoteChainID = flag.Int("remote-chainid", getEnvAsIntOrDefault("REMOTE_CHAINID", 45206), "chain id of remote chain")
	amount        = flag.Int("amount", getEnvAsIntOrDefault("AMOUNT", 1), "amount of token to transfer")

	privKey = flag.String("priv-key", getEnvAsStrOrDefault("PRIV_KEY", "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"), "priv key of wallet without 0x prefix")
)

func main() {
	flag.Parse()

	fbRpc1 := flashbotsrpc.NewFlashbotsRPC(*gethUrlOrigin)
	fbRpc2 := flashbotsrpc.NewFlashbotsRPC(*gethUrlRemote)
	fbRpcJavelin := flashbotsrpc.NewFlashbotsRPC(*javelinUrl) // the javelin-rpc endpoint, which caches bundles and simulate them and send them in a preset interval(500ms)

	pkHex := strings.TrimLeft(*privKey, "0x")
	privKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := getAddressFromPriv(privKey)
	if err != nil {
		log.Fatal(err)
	}

	originChainID := big.NewInt(int64(*originChainID))
	remoteChainID := big.NewInt(int64(*remoteChainID))

	// populate two transactions, one on origin chain(45206) and one on remote chain(45207), those transactions will be included in one bundle
	// once the bundle get picked, javelin-rpc will simulate the picked bundles and filtered out the failed ones and send the successful ones to be executed atomically
	txOrigin, err := genTransfer(fbRpc1, originChainID, privKey, addr, *amount, 21000, nil)
	if err != nil {
		log.Fatal(err)
	}
	txOriginRaw, err := txOrigin.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	txRemote, err := genTransfer(fbRpc2, remoteChainID, privKey, addr, *amount, 21000, nil)
	if err != nil {
		log.Fatal(err)
	}
	txRemoteRaw, err := txRemote.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}

	txs := make(map[string]string)
	txs[hexutil.EncodeBig(originChainID)] = hexutil.Encode(txOriginRaw)
	txs[hexutil.EncodeBig(remoteChainID)] = hexutil.Encode(txRemoteRaw)

	// unused for javelin-rpc since currently the atomic execution is backed by Nodekit SEQ, the bundle will be included for sure in a SEQ block
	// but txns will be included on uncertain blocks for each rollups, which is expected as we assure the atomicity by SEQ
	blockNumber := make(map[string]string)
	blockNumber[hexutil.EncodeBig(originChainID)] = "0x0"
	blockNumber[hexutil.EncodeBig(remoteChainID)] = "0x0"

	bundleArgs := flashbotsrpc.FlashbotsSendBundleCrossRollupRequest{
		Txs:         txs,
		BlockNumber: blockNumber,
	}

	// send bundle
	resp, err := fbRpcJavelin.FlashbotsSendBundleCrossRollup(privKey, bundleArgs)
	if err != nil {
		log.Fatal(err)
	}
	bundleHash := resp.BundleHash
	log.Printf("submitted bundle %s\n", bundleHash)

	// query bundle status
	accepted := false
	for i := 0; i < 10; i++ {
		if accepted {
			break
		}
		time.Sleep(500 * time.Millisecond)
		bundleStatusResp, err := fbRpcJavelin.FlashbotsGetBundleCrossRollupStatus(privKey, flashbotsrpc.FlashbotsGetBundleCrossRollupStatusRequest{
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
			accepted = true
			log.Printf("bundle accepted")
		case 0x1:
			log.Fatalf("bundle rejected, %+v", bundleStatusResp.Status)
		case 0x2:
			log.Printf("bundle not sent to SEQ yet, Status: %s", bundleStatusResp.Status)
			continue
		}
	}

	// wait tx to be included on both chain
	var wg sync.WaitGroup
	wg.Add(2)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	go func() {
		defer wg.Done()
		waitTx(ctx, fbRpc1, txOrigin.Hash().Hex())
	}()

	go func() {
		defer wg.Done()
		waitTx(ctx, fbRpc2, txRemote.Hash().Hex())
	}()

	wg.Wait()
}

func waitTx(ctx context.Context, rpc *flashbotsrpc.FlashbotsRPC, txHash string) {
	log.Printf("waiting tx: %s\n", txHash)
	for {
		select {
		case <-ctx.Done():
			log.Printf("context cancled: %s\n", ctx.Err())
			return
		case <-time.After(1 * time.Second):
			receipt, err := rpc.EthGetTransactionReceipt(txHash)
			if err != nil || receipt.TransactionHash == "" {
				log.Printf("unable to get receipt: err(%+v) or receipt is null\n", err)
			} else {
				log.Printf("got receipt: %+v\n", receipt)
				return
			}
		}
	}
}

func genTransfer(rpc *flashbotsrpc.FlashbotsRPC, chainID *big.Int, privKey *ecdsa.PrivateKey, toAddress *common.Address, amount int, gasLimit uint64, data []byte) (*ethtypes.Transaction, error) {
	gasPrice, err := rpc.EthGasPrice()
	if err != nil {
		return nil, err
	}
	value := big.NewInt(int64(amount))
	pubKey := privKey.Public()
	pubkeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	sender := crypto.PubkeyToAddress(*pubkeyECDSA)
	nonce, err := rpc.EthGetTransactionCount(sender.Hex(), "latest")
	if err != nil {
		return nil, err
	}

	tx := ethtypes.NewTransaction(uint64(nonce), *toAddress, value, gasLimit, &gasPrice, nil)
	return ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), privKey)
}

func getAddressFromPriv(privKey *ecdsa.PrivateKey) (*common.Address, error) {
	pubkey := privKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("unable to reflect pubkey to ECDSA")
	}
	addr := crypto.PubkeyToAddress(*pubkeyECDSA)
	return &addr, nil
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
