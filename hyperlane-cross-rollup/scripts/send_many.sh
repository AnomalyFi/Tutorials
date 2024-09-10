#!/usr/bin/bash

GETH_ORIGIN=http://127.0.0.1:9095
GETH_REMOTE=http://127.0.0.1:9096
ORIGIN_CHAINID=45205
REMOTE_CHAINID=45206
KEY1=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
KEY2=7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY1 &
go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY2 &

wait