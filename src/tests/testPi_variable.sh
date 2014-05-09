#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

NUM_CLIENTS=10

for (( i = 0 ; i < $NUM_CLIENTS ; i++ ))
do
    go run $CLIENT_GO -b=$TESTS/pure_basic_25_"$i".txt > $TESTS/logs/client"$i"_pb25_10.log &
    CLIENT_PID[$i]=$!
done

for (( i = 0 ; i < $NUM_CLIENTS ; i++ ))
do
    echo "[TEST] Waiting for client $i to finish"
    wait ${CLIENT_PID[$i]}
done

T="$(($(date +%s%N)-$T))"
M="$((T/1000000))"
echo "[TEST] Time Elapsed in milliseconds: ${M}"

for (( i = 0 ; i < $NUM_CLIENTS ; i++ ))
do
    echo "$(tail -n 1 logs/client"$i"_pb25_10.log | awk 'NF>1{print $NF}')"
done

echo "[TEST] Correctness not yet checked"
echo "[TEST] Finished."

