#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

go run $CLIENT_GO -b=$TESTS/pure_basic_25_1.txt > $TESTS/logs/client1_pb25_10.log &
CLIENT_PID1=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_2.txt > $TESTS/logs/client2_pb25_10.log &
CLIENT_PID2=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_3.txt > $TESTS/logs/client3_pb25_10.log &
CLIENT_PID3=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_4.txt > $TESTS/logs/client4_pb25_10.log &
CLIENT_PID4=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_5.txt > $TESTS/logs/client5_pb25_10.log &
CLIENT_PID5=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_6.txt > $TESTS/logs/client6_pb25_10.log &
CLIENT_PID6=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_7.txt > $TESTS/logs/client7_pb25_10.log &
CLIENT_PID7=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_8.txt > $TESTS/logs/client8_pb25_10.log &
CLIENT_PID8=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_9.txt > $TESTS/logs/client9_pb25_10.log &
CLIENT_PID9=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_A.txt > $TESTS/logs/clientA_pb25_10.log &
CLIENT_PID10=$!

echo "[TEST] Waiting for client 1 to finish"
wait $CLIENT_PID1
echo "[TEST] Waiting for client 2 to finish"
wait $CLIENT_PID2
echo "[TEST] Waiting for client 3 to finish"
wait $CLIENT_PID3
echo "[TEST] Waiting for client 4 to finish"
wait $CLIENT_PID4
echo "[TEST] Waiting for client 5 to finish"
wait $CLIENT_PID5
echo "[TEST] Waiting for client 6 to finish"
wait $CLIENT_PID6
echo "[TEST] Waiting for client 7 to finish"
wait $CLIENT_PID7
echo "[TEST] Waiting for client 8 to finish"
wait $CLIENT_PID8
echo "[TEST] Waiting for client 9 to finish"
wait $CLIENT_PID9
echo "[TEST] Waiting for client A to finish"
wait $CLIENT_PID10

T="$(($(date +%s%N)-$T))"
M="$((T/1000000))"
echo "[TEST] Time Elapsed in milliseconds: ${M}"

client1Time="$(tail -n 1 logs/client1_pb25_10.log | awk 'NF>1{print $NF}')"
client2Time="$(tail -n 1 logs/client2_pb25_10.log | awk 'NF>1{print $NF}')"
client3Time="$(tail -n 1 logs/client3_pb25_10.log | awk 'NF>1{print $NF}')"
client4Time="$(tail -n 1 logs/client4_pb25_10.log | awk 'NF>1{print $NF}')"
client5Time="$(tail -n 1 logs/client5_pb25_10.log | awk 'NF>1{print $NF}')"
client6Time="$(tail -n 1 logs/client6_pb25_10.log | awk 'NF>1{print $NF}')"
client7Time="$(tail -n 1 logs/client7_pb25_10.log | awk 'NF>1{print $NF}')"
client8Time="$(tail -n 1 logs/client8_pb25_10.log | awk 'NF>1{print $NF}')"
client9Time="$(tail -n 1 logs/client9_pb25_10.log | awk 'NF>1{print $NF}')"
clientATime="$(tail -n 1 logs/clientA_pb25_10.log | awk 'NF>1{print $NF}')"
echo $client1Time
echo $client2Time
echo $client3Time
echo $client4Time
echo $client5Time
echo $client6Time
echo $client7Time
echo $client8Time
echo $client9Time
echo $clientATime

echo "[TEST] Correctness not yet checked"
echo "[TEST] Finished."

