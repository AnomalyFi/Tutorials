#!/usr/bin/bash

ROOT=$(pwd)
INC=${INC:-0}
HOST=${HOST:-127.0.0.1}
SPAM_ITER=${SPAM_ITER:-20}

LZ_PROGRAM_DIR="./layerzero-cross-rollup"
HYP_PROGRAM_DIR="./hyperlane-cross-rollup"

JAVELIN_RPC=${JAVELIN_RPC:-http://127.0.0.1:3000/rpc}

ORIGIN_CHAINID=$((45200 + $INC))
REMOTE_CHAINID=$((45200 + $INC + 1))
ORIGIN_GETH_PORT=$((19545 + $INC))
REMOTE_GETH_PORT=$((19545 + $INC + 1))

ORIGIN_GETH="http://$HOST:$ORIGIN_GETH_PORT"
REMOTE_GETH="http://$HOST:$REMOTE_GETH_PORT"

MAILBOX_ORIGIN=${MAILBOX_ORIGIN:-0xa513E6E4b8f2a923D98304ec87F64353C4D5C853}
MAILBOX_REMOTE=${MAILBOX_REMOTE:-0xa513E6E4b8f2a923D98304ec87F64353C4D5C853}
TR_ORIGIN=${TR_ORIGIN:-0x4A679253410272dd5232B3Ff7cF5dbB88f295319}
TR_REMOTE=${TR_REMOTE:-0x4A679253410272dd5232B3Ff7cF5dbB88f295319}

OFT_ORIGIN=${OFT_ORIGIN:-0x985060F8b809F08392FB4E23622E9E6881c22d0b}
OFT_REMOTE=${OFT_REMOTE:-0x985060F8b809F08392FB4E23622E9E6881c22d0b}

echo "origin chain: $ORIGIN_CHAINID; geth url: $ORIGIN_GETH"
echo "remote chain: $REMOTE_CHAINID; geth url: $REMOTE_GETH"
echo "javelin rpc: $JAVELIN_RPC"

PRIVATE_KEYS=("0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6" "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
# the first two addresses are the addresses of the above private keys
TARGET_ADDRS=("0x90F79bf6EB2c4f870365E785982E1f101E93b906" "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC" "0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

NONCES_ORIGIN=()
NONCES_REMOTE=()

# Loop through each address and get the nonce
for addr in "${TARGET_ADDRS[@]}"; do
    echo "checking nonce of $addr"
    # Fetch the nonce using cast nonce and store it in a variable
    nonce_origin=$(cast nonce "$addr" --rpc-url $ORIGIN_GETH)
    nonce_remote=$(cast nonce "$addr" --rpc-url $REMOTE_GETH)
    
    # Append the nonce to the NONCES array
    NONCES_ORIGIN+=("$nonce_origin")
    NONCES_REMOTE+=("$nonce_remote")
done


# Print the addresses and their corresponding nonces
for i in "${!TARGET_ADDRS[@]}"; do
    echo "Address: ${TARGET_ADDRS[i]}, Nonce Origin: ${NONCES_ORIGIN[i]}, Nonce Remote: ${NONCES_REMOTE[i]}"
done

REVERSE_DIRECTION="reversed"
NORMAL_DIRECTION="normal"

function send_hyperlane_bundle() {
    local priv_key=$1
    local direction=$2 # normal or reversed 
    local target=$3
    local nonce=$4
    local timestamp=$5

    echo "sending hyperlane bundle with privKey: $priv_key, direction: $direction, nonce inc: $nonce"

    cd $HYP_PROGRAM_DIR
    if test "$direction" = "$NORMAL_DIRECTION" 
    then
        go run main.go --javelin $JAVELIN_RPC --geth-origin $ORIGIN_GETH --geth-remote $REMOTE_GETH --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $priv_key --amount 1 --nonce $nonce --target $target --timestamp $timestamp &
    else
        go run main.go --javelin $JAVELIN_RPC --geth-origin $REMOTE_GETH --geth-remote $ORIGIN_GETH --origin-chainid $REMOTE_CHAINID --remote-chainid $ORIGIN_CHAINID --priv-key $priv_key --amount 1 --nonce $nonce --target $target --timestamp $timestamp &
    fi

    cd $ROOT
    PID=$!
    return $PID
}

function send_layerzero_bundle() {
    local priv_key=$1
    local direction=$2 # normal or reversed 
    local target=$3
    local nonce=$4
    local timestamp=$5

    echo "sending layerzero bundle with privKey: $priv_key, direction: $direction, nonc: $nonce"

    cd $LZ_PROGRAM_DIR
    if test "$direction" = "$NORMAL_DIRECTION" 
    then
        go run main.go --javelin $JAVELIN_RPC --geth-origin $ORIGIN_GETH --geth-remote $REMOTE_GETH --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --priv-key $priv_key --amount 1 --nonce $nonce --target $target --timestamp $timestamp &
    else
        go run main.go --javelin $JAVELIN_RPC --geth-origin $REMOTE_GETH --geth-remote $ORIGIN_GETH --origin-chainid $REMOTE_CHAINID --remote-chainid $ORIGIN_CHAINID --priv-key $priv_key --amount 1 --nonce $nonce --target $target --timestamp $timestamp &
    fi

    cd $ROOT
    PID=$!
    return $PID
}

function send_transfer() {
    local target_chain=$1
    local privkey=$2
    local nonce=$3
    local to=$4

    echo "sending transfer from $privkey to $to with nonce: $nonce on chain $target_chain"

    cast send $to --rpc-url $target_chain --private-key $privkey --nonce $nonce &
    PID=$!
    return $PID
}


# send_hyperlane_bundle $PRIVATE_KEY $NORMAL_DIRECTION 0
# PID=$!
# wait $PID
# echo "status of process: $?"

GLOBAL_TIMESTAMP=$(date +%s%3N)

function increment_timestamp() {
    GLOBAL_TIMESTAMP=$((GLOBAL_TIMESTAMP+1))
}


# send_layerzero_bundle ${PRIVATE_KEYS[0]} $NORMAL_DIRECTION ${TARGET_ADDRS[0]} 0
# PID=$!
# wait $PID
# echo "status of process: $?"

# increment_timestamp

# send_layerzero_bundle ${PRIVATE_KEYS[0]} $NORMAL_DIRECTION ${TARGET_ADDRS[0]} 0
# PID=$!
# wait $PID
# echo "status of process: $?"

# send_transfer $REMOTE_GETH ${PRIVATE_KEYS[1]} 1 ${TARGET_ADDRS[1]}

# increment_timestamp
# refresh current nonce
NONCES_ORIGIN=()
NONCES_REMOTE=()

# Loop through each address and get the nonce
for addr in "${TARGET_ADDRS[@]}"; do
    # Fetch the nonce using cast nonce and store it in a variable
    nonce_origin=$(cast nonce "$addr" --rpc-url $ORIGIN_GETH)
    nonce_remote=$(cast nonce "$addr" --rpc-url $REMOTE_GETH)
    
    # Append the nonce to the NONCES array
    NONCES_ORIGIN+=("$nonce_origin")
    NONCES_REMOTE+=("$nonce_remote")
done

# Print the addresses and their corresponding nonces
for i in "${!TARGET_ADDRS[@]}"; do
    echo "Address: ${TARGET_ADDRS[i]}, Nonce Origin: ${NONCES_ORIGIN[i]}, Nonce Remote: ${NONCES_REMOTE[i]}"
done

nonce_origin=${NONCES_ORIGIN[0]}
send_transfer $ORIGIN_GETH ${PRIVATE_KEYS[0]} $((nonce_origin)) ${TARGET_ADDRS[1]}
TX_PID=$!
wait $TX_PID

send_layerzero_bundle ${PRIVATE_KEYS[0]} $NORMAL_DIRECTION ${TARGET_ADDRS[0]} $((nonce_origin+1)) $GLOBAL_TIMESTAMP
LZ_PID=$!
wait $LZ_PID
increment_timestamp

send_hyperlane_bundle ${PRIVATE_KEYS[0]} $NORMAL_DIRECTION ${TARGET_ADDRS[0]} $((nonce_origin+2)) $GLOBAL_TIMESTAMP
HYP_PID=$!
wait $HYP_PID
increment_timestamp