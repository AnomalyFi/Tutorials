## Javelin Cross Rollup Bundle Example

Before we start, please refers to the bundle API specification at [Bundle API](https://nodekit.gitbook.io/nodekit-documentation/builders-guide/javelin/bundle-api). The up to update `eth_sendBundleCrossRollup` JSON-RPC request looks like this

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "eth_sendBundleCrossRollup",
  "params": [
    {
      txs,               // HashMap[String, String], This contains a mapping of signed transactions alonside their corresponding chainIds which we want to execute in an atomic bundle.
      blockNumbers,       // HashMap[String, String], This contains a mapping of chainIds alongside the corresponding hex encoded block number which this bundle will be valid on.  
      minTimestamp,      // (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
      maxTimestamp,      // (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
      revertingTxHashes, // (Optional) Array[String], A list of tx hashes that are allowed to revert
      replacementUuid,   // (Optional) String, UUID that can be used to cancel/replace this bundle
    }
  ]
}
```

Where the transactions are stored at `txs` fields in a map, the key of the map is hex encoded chainID, e.g. 100 -> "0x64", and the value is one transaction for that domain. Currently only one transaction per domain is supported and in the future we may support more transactions for one domain. 

### FlashbotsRPC in NodeKit

Since the Javelin Superbuilder was originally the `Flashbots Proteact RPC Endpoint`, based on the original flashbots-rpc implementation, we forked off and made [our client](https://github.com/AnomalyFi/flashbotsrpc) for cross-rollup bundle support. In this tutorial, the following methods are used

```go
fbRPC.FlashbotsSendBundleCrossRollup(privKey, bundleArgs) // send bundle
fbRpcJavelin.FlashbotsGetBundleCrossRollupStatus(privKey, statusArgs) // query bundle status
```

For each of them the arguments and the responses are

```go
type FlashbotsSendBundleCrossRollupRequest struct {
	Txs               map[string]string `json:"txs"`
	BlockNumber       map[string]string `json:"blockNumber"`
	MinTimestamp      *uint64           `json:"minTimestamp,omitempty"` // (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
	MaxTimestamp      *uint64           `json:"maxTimestamp,omitempty"` // (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
	RevertingTxHashes []string          `json:"revertingTxHashes,omitempty"`
	ReplacementUuid   string            `json:"replacementUUID,omitempty"`
}
type FlashbotsSendBundleCrossRollupResponse struct {
	BundleHash string `json:"bundleHash"`
}
```

```go
type FlashbotsGetBundleCrossRollupStatusRequest struct {
	BundleHash string `json:"bundleHash"`
}
type FlashbotsGetBundleCrossRollupStatusResponse struct {
	StatusCode int    `json:"statusCode"` // 0x0 -> included; 0x1 -> rejected; 0x2 sent but not included;
	Status     string `json:"status"`     // descriptive string
}
```

So the general steps to send a cross rollup bundle will be

1. Construct cross rollup bundle & send
2. Query bundle status
3. If bundle is accepted, query status of each transactions since by accepting it only means the bundle is accepted by Javelin RPC and Arcadia, but we still need to wait on the transactions land on corresponding chains

### Usage

```shell
➜  cross-rollup git:(main) ✗ go run main.go --help
Usage of /tmp/go-build1796757564/b001/exe/main:
  -amount int
        amount of token to transfer (default 1)
  -geth-origin string
        geth rpc url on origin chain (default "http://127.0.0.1:19550")
  -geth-remote string
        geth rpc url on remote chain (default "http://127.0.0.1:19551")
  -javelin string
        rpc url of javelin rpc (default "http://127.0.0.1:3000/rpc")
  -origin-chainid int
        chain id of origin chain (default 45200)
  -priv-key string
        priv key of wallet without 0x prefix (default "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
  -remote-chainid int
        chain id of remote chain (default 45206)
```

