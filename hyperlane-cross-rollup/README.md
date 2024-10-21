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
➜  hyperlane-cross-rollup git:(main) ✗ go run main.go --help
Usage of /tmp/go-build1293359915/b001/exe/main:
  -amount int
        amount of token to transfer (default 1)
  -geth-origin string
        geth rpc url on origin chain (default "http://127.0.0.1:9090")
  -geth-remote string
        geth rpc url on remote chain (default "http://127.0.0.1:9091")
  -javelin string
        rpc url of javelin rpc (default "http://127.0.0.1:3000/rpc")
  -mailbox-origin string
        mailbox address on origin chain (default "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853")
  -mailbox-remote string
        mailbox address on remote chain (default "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853")
  -nonce-inc int
        the nonce adjustment applied to the latest nonce
  -origin-chainid int
        chain id of origin chain (default 45200)
  -priv-key string
        priv key of wallet without 0x prefix (default "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
  -remote-chainid int
        chain id of remote chain (default 45201)
  -router-origin string
        token router address on origin chain (default "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
  -router-remote string
        token router address on remote chain (default "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
  -token-type string
        token type on origin chain to transfer from origin chain to remote chain (default "native")
```

