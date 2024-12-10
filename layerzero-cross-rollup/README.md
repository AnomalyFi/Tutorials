## Layerzero atomic cross rollup bundle example

In this example, we implemented the minimal codes to send Layerzero OFTs between origin chain and target chain atomically through Javelin Cross-Rollup Bundle. For detail, check codes at `main.go` 

### Usage

To run the go program to send hyperlane cross rollup bundle, simply export the following env variables

| Env Name       | Value                                 | Description                                                  |
| -------------- | ------------------------------------- | ------------------------------------------------------------ |
| JAVELIN        | http://devnet.nodekit.xyz/javelin/rpc | the rpc endpoint of javelin-rpc                              |
| GETH_ORIGIN    | http://devnet.nodekit.xyz/r5          | the geth endpoint to the origin where tokens will be transferred from |
| GETH_REMOTE    | http://devnet.nodekit.xyz/r5          | the geth endpoint to the destination where tokens will be transferred to |
| ORIGIN_CHAINID | 45205                                 |                                                              |
| REMOTE_CHAINID | 45206                                 |                                                              |
| KEY            | \<you-private-key-in-hex\>            | Your private key in hex with 0x prefix                       |

```shell
go run main.go --javelin $JAVELIN --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY --amount 1
```



Program usage:

```shell
➜  layerzero-cross-rollup git:(main) ✗ go run main.go --help
Usage of /tmp/go-build2295665594/b001/exe/main:
  -amount int
        amount of oft to transfer (default 1000000000000)
  -geth-origin string
        geth rpc url on origin chain (default "http://127.0.0.1:9090")
  -geth-remote string
        geth rpc url on remote chain (default "http://127.0.0.1:9091")
  -javelin string
        rpc url of javelin rpc (default "http://127.0.0.1:3000/rpc")
  -nonce-inc int
        the nonce adjustment applied to the latest nonce
  -oft-origin string
        oft address on origin chain (default "0x985060F8b809F08392FB4E23622E9E6881c22d0b")
  -oft-remote string
        oft address on remote chain (default "0x985060F8b809F08392FB4E23622E9E6881c22d0b")
  -origin-chainid int
        chain id of origin chain (default 45205)
  -priv-key string
        priv key of wallet without 0x prefix (default "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
  -remote-chainid int
        chain id of remote chain (default 45206)
```

