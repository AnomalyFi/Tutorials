package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

const R1 = "http://devnet.nodekit.xyz/javelin/rpc"
const R2 = "http://devnet.nodekit.xyz/javelin-reverse/rpc"
const JAVELIN = "http://devnet.nodekit.xyz/javelin/rpc"
const ORIGIN_CHAINID = 45206
const REMOTE_CHAINID = 45207

func main() {
	fbRpc1 := flashbotsrpc.NewFlashbotsRPC(R1)
	fbRpc2 := flashbotsrpc.NewFlashbotsRPC(R2)
	fbRpcJavelin := flashbotsrpc.NewFlashbotsRPC(JAVELIN) // the javelin-rpc endpoint, which caches bundles and simulate them and send them in a preset interval(500ms)

	// the wallet private key to issue cross-rollup bundle
	testKeyStr := "47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a"
	privKey, err := crypto.HexToECDSA(testKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := getAddressFromPriv(privKey)
	if err != nil {
		log.Fatal(err)
	}

	originChainID := big.NewInt(ORIGIN_CHAINID)
	remoteChainID := big.NewInt(REMOTE_CHAINID)

	// populate two transactions, one on origin chain(45206) and one on remote chain(45207), those transactions will be included in one bundle
	// once the bundle get picked, javelin-rpc will simulate the picked bundles and filtered out the failed ones and send the successful ones to be executed atomically
	txOrigin, err := genTransfer(fbRpc1, originChainID, privKey, addr, 20, 21000, nil)
	if err != nil {
		log.Fatal(err)
	}
	txOriginRaw, err := txOrigin.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	txRemote, err := genTransfer(fbRpc2, remoteChainID, privKey, addr, 50, 21000, nil)
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
