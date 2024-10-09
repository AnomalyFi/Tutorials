#!/usr/bin/bash
PWD=$(pwd)
INC=${INC:-0}
HOST=${HOST:-127.0.0.1}
JAVELIN_RPC=${JAVELIN_RPC:-http://$HOST:3000/rpc}
AMOUNT=${AMOUNT:-1}
PRIVATE_KEY=${PRIVATE_KEY:-0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6}
PRIVATE_KEY2=${PRIVATE_KEY2:-0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a}
PRIVATE_KEY3=${PRIVATE_KEY3:-0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a}
ORIGIN_CHAINID=$((45200 + $INC))
REMOTE_CHAINID=$((45200 + $INC + 1))
ORIGIN_GETH_PORT=$((19545 + $INC))
REMOTE_GETH_PORT=$((19545 + $INC + 1))

ORIGIN_GETH="http://$HOST:$ORIGIN_GETH_PORT"
REMOTE_GETH="http://$HOST:$REMOTE_GETH_PORT"

echo "origin chain: $ORIGIN_CHAINID; geth url: $ORIGIN_GETH"
echo "remote chain: $REMOTE_CHAINID; geth url: $REMOTE_GETH"
echo "javelin rpc: $JAVELIN_RPC"

# to hold multiple PIDs
PIDS=()

# 1st tx
# gas too low so fails which is what we want
cast send 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --rpc-url $ORIGIN_GETH --private-key $PRIVATE_KEY --value $AMOUNT --gas-limit 20990 --gas-price 10000 &
PIDS+=($!)

# 2nd tx
# min gas limit
cast send 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC --rpc-url $ORIGIN_GETH --private-key $PRIVATE_KEY2 --value $ANOTHER_AMOUNT --gas-limit 21000 --gas-price 100000 --nonce 39 &
PIDS+=($!)

# 3rd tx
# high gas enough to exceed gas limit
cast send 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 --rpc-url $ORIGIN_GETH --private-key $PRIVATE_KEY3 --value $YET_ANOTHER_AMOUNT --gas-limit 1000000 --gas-price 100000000000 --nonce 39 &
PIDS+=($!)

# wait for PIDs to complete
for PID in "${PIDS[@]}"; do
    wait $PID
done

cd cross-rollup

go run main.go --javelin $JAVELIN_RPC --geth-origin $ORIGIN_GETH --origin-chainid $ORIGIN_CHAINID --amount $AMOUNT &
PID_GO=$!

wait $PID_GO
