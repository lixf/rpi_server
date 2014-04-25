#!/bin/bash

if [ -z $GOPATH ]; then
    echo "FAIL: GOPATH environment variable not set"
    exit 1
fi

echo "Basic Test: Two workers, one master, three clients."

WORKER_PORT1=$(((RANDOM % 10000) + 10000))
WORKER_PORT2=$(((RANDOM % 10000) + 10000))
WORKER_GO=$GOPATH/src/runners/wrunner/wrunner.go
MASTER_GO=$GOPATH/src/runners/mrunner/mrunner.go
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TEST_FILE1=$GOPATH/src/tests/longGetPost1.txt
TEST_FILE2=$GOPATH/src/tests/longGetPost2.txt
TEST_FILE3=$GOPATH/src/tests/longGetPost3.txt

go run $WORKER_GO -port=$WORKER_PORT1 > "logs/worker1.log" &
WORKER1_PID=$!

go run $WORKER_GO -port=$WORKER_PORT2 > "logs/worker2.log" &
WORKER2_PID=$!

go run $MASTER_GO -N=2 > "logs/master.log" &
MASTER_PID=$!

sleep 3
echo "Starting timer:"
T="$(date +%s%N)"

# Do some work here
go run $CLIENT_GO -b=$TEST_FILE1 > "logs/client1.log" &
CLIENT_PID1=$!
go run $CLIENT_GO -b=$TEST_FILE2 > "logs/client2.log" &
CLIENT_PID2=$!
go run $CLIENT_GO -b=$TEST_FILE3 > "logs/client3.log" &
CLIENT_PID3=$!

echo "[TEST] Waiting for clients to finish"
wait $CLIENT_PID1
echo "[TEST] Client 1 finished."
wait $CLIENT_PID2
echo "[TEST] Client 2 finished."
wait $CLIENT_PID3
echo "[TEST] Client 3 finished."

T="$(($(date +%s%N)-$T))"
M="$((T/1000000))"
echo "[TEST] Time Elapsed in milliseconds: ${M}"
echo "[TEST] Correctness not yet checked"
echo "[TEST] Finish. Control + C to kill master and workers."

wait $WORKER1_PID
wait $WORKER2_PID
wait $MASTER_PID
