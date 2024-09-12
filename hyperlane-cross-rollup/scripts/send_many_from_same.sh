#!/usr/bin/bash

GETH_ORIGIN=http://127.0.0.1:9095
GETH_REMOTE=http://127.0.0.1:9096
ORIGIN_CHAINID=45205
REMOTE_CHAINID=45206
KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY --nonce-inc 0 --amount 1 &
sleep .1
go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $KEY --nonce-inc 1 --amount 2 & 

wait
