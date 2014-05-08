#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

go run $CLIENT_GO -b=$TESTS/pure_basic_1000.txt > $TESTS/logs/client1_pb1000_3.log &
CLIENT_PID1=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_1000.txt > $TESTS/logs/client2_pb1000_3.log &
CLIENT_PID2=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_1000.txt > $TESTS/logs/client3_pb1000_3.log &
CLIENT_PID3=$!

echo "[TEST] Waiting for client 1 to finish"
wait $CLIENT_PID1
echo "[TEST] Waiting for client 2 to finish"
wait $CLIENT_PID2
echo "[TEST] Waiting for client 3 to finish"
wait $CLIENT_PID3

T="$(($(date +%s%N)-$T))"
M="$((T/1000000))"
echo "[TEST] Time Elapsed in milliseconds: ${M}"
echo "[TEST] Correctness not yet checked"
echo "[TEST] Finished."

