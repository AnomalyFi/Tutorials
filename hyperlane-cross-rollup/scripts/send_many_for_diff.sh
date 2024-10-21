#!/usr/bin/bash

GETH_ORIGIN=http://devnet.nodekit.xyz/r5
GETH_REMOTE=http://devnet.nodekit.xyz/r6
JAVELIN=http://devnet.nodekit.xyz/javelin/rpc
ORIGIN_CHAINID=45205
REMOTE_CHAINID=45206
KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

go run main.go --javelin $JAVELIN --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY --amount 1 &
sleep .1
go run main.go --javelin $JAVELIN --geth-origin $GETH_REMOTE --geth-remote $GETH_ORIGIN --origin-chainid $REMOTE_CHAINID --remote-chainid $ORIGIN_CHAINID --priv-key $KEY --amount 2 --router-origin 0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E --router-remote 0x67d269191c92Caf3cD7723F116c85e6E9bf55933 & 

wait
