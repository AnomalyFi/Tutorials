#!/usr/bin/sh

GETH_ORIGIN=${GETH_ORIGIN:-http://127.0.0.1:19550}
GETH_REMOTE=${GETH_REMOTE:-http://127.0.0.1:19551}
ORIGIN_CHAINID=${ORIGIN_CHAINID:-45205}
REMOTE_CHAINID=${REMOTE_CHAINID:-45206}
OFT_ORIGIN=${OFT_ORIGIN:-0xfc2a06C490A7Fc32EF83A9AaE8554d522a7AE213}
OFT_REMOTE=${OFT_REMOTE:-0xacF1571337efbf93b3E3Cf037EdcC1ef715eC3DC}
KEY=${KEY:-0xacF1571337efbf93b3E3Cf037EdcC1ef715eC3DC}

echo $OFT_ORIGIN

go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --oft-origin $OFT_ORIGIN --oft-remote $OFT_REMOTE &
PID1=$!
sleep .1
go run main.go --geth-origin $GETH_ORIGIN --geth-remote $GETH_REMOTE --origin-chainid $ORIGIN_CHAINID --remote-chainid $REMOTE_CHAINID --oft-origin $OFT_ORIGIN --oft-remote $OFT_REMOTE --nonce-inc 1 &
PID2=$!

wait $PID1
wait $PID2