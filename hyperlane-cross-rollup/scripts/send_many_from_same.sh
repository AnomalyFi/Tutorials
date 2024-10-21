#!/usr/bin/bash

GETH_ORIGIN=http://devnet.nodekit.xyz/r5
GETH_REMOTE=http://devnet.nodekit.xyz/r6
JAVELIN=http://devnet.nodekit.xyz/javelin/rpc
ORIGIN_CHAINID=45205
REMOTE_CHAINID=45206
KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
KEY2=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

go run main.go --javelin $JAVELIN --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY --nonce-inc 0 --amount 1 &
sleep .1
go run main.go --javelin $JAVELIN --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY2 --nonce-inc 0 --amount 2 & 

wait
