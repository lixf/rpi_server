#!/bin/bash

if [ -z $GOPATH ]; then
    echo "FAIL: GOPATH environment variable not set"
    exit 1
fi

echo "Basic Test: Two workers, one master, one client."

WORKER_PORT1=$(((RANDOM % 10000) + 10000))
WORKER_PORT2=$(((RANDOM % 10000) + 10000))
WORKER_GO=$GOPATH/src/runners/wrunner/wrunner.go
MASTER_GO=$GOPATH/src/runners/mrunner/mrunner.go
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TEST_FILE=basicGetPost.txt

go run $WORKER_GO -port=$WORKER_PORT1 > "logs/worker1.log" &
WORKER1_PID=$!

go run $WORKER_GO -port=$WORKER_PORT2 > "logs/worker2.log" &
WORKER2_PID=$!

go run $MASTER_GO -N=2 > "logs/master.log" &
MASTER_PID=$!

sleep 3
go run $CLIENT_GO -b=$TEST_FILE > "logs/client.log" &
sleep 4


cat logs/master.log

echo "[TEST] Finish. Control + C to kill master and workers."

wait $WORKER1_PID
wait $WORKER2_PID
wait $MASTER_PID
