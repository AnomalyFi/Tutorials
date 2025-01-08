#!/usr/bin/bash
PWD=$(pwd)
INC=${INC:-0}
HOST=${HOST:-127.0.0.1}
JAVELIN_RPC=${JAVELIN_RPC:-http://127.0.0.1:3000/rpc}
# the following two should be different to guarantee the hash of them are different
AMOUNT=${AMOUNT:-1}
AMOUNT_GETH=${AMOUNT_GETH:-2}
PRIVATE_KEY=${PRIVATE_KEY:-0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6}
ORIGIN_CHAINID=$((45200 + $INC))
REMOTE_CHAINID=$((45200 + $INC + 1))
ORIGIN_GETH_PORT=$((19545 + $INC))
REMOTE_GETH_PORT=$((19545 + $INC + 1))

ORIGIN_GETH="http://$HOST:$ORIGIN_GETH_PORT"
REMOTE_GETH="http://$HOST:$REMOTE_GETH_PORT"

echo "origin chain: $ORIGIN_CHAINID; geth url: $ORIGIN_GETH"
# echo "remote chain: $REMOTE_CHAINID; geth url: $REMOTE_GETH"
echo "javelin rpc: $JAVELIN_RPC"

cd cross-rollup-single
go run main.go --javelin $JAVELIN_RPC --geth-origin $ORIGIN_GETH --origin-chainid $ORIGIN_CHAINID --amount $AMOUNT &
PID_GO=$!

cast send 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --rpc-url $ORIGIN_GETH --private-key $PRIVATE_KEY --value $AMOUNT_GETH &
PID_CAST=$!

wait $PID_CAST
wait $PID_GO

