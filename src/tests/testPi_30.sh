#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

go run $CLIENT_GO -b=$TESTS/pure_basic_25_1.txt > $TESTS/logs/client01_pb25_20.log &
CLIENT_PID01=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_2.txt > $TESTS/logs/client02_pb25_20.log &
CLIENT_PID02=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_3.txt > $TESTS/logs/client03_pb25_20.log &
CLIENT_PID03=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_4.txt > $TESTS/logs/client04_pb25_20.log &
CLIENT_PID04=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_5.txt > $TESTS/logs/client05_pb25_20.log &
CLIENT_PID05=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_6.txt > $TESTS/logs/client06_pb25_20.log &
CLIENT_PID06=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_7.txt > $TESTS/logs/client07_pb25_20.log &
CLIENT_PID07=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_8.txt > $TESTS/logs/client08_pb25_20.log &
CLIENT_PID08=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_9.txt > $TESTS/logs/client09_pb25_20.log &
CLIENT_PID09=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_10.txt > $TESTS/logs/client0x_pb25_20.log &
CLIENT_PID10=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_1.txt > $TESTS/logs/client11_pb25_20.log &
CLIENT_PID11=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_2.txt > $TESTS/logs/client12_pb25_20.log &
CLIENT_PID12=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_3.txt > $TESTS/logs/client13_pb25_20.log &
CLIENT_PID13=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_4.txt > $TESTS/logs/client14_pb25_20.log &
CLIENT_PID14=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_5.txt > $TESTS/logs/client15_pb25_20.log &
CLIENT_PID15=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_6.txt > $TESTS/logs/client16_pb25_20.log &
CLIENT_PID16=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_7.txt > $TESTS/logs/client17_pb25_20.log &
CLIENT_PID17=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_8.txt > $TESTS/logs/client18_pb25_20.log &
CLIENT_PID18=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_9.txt > $TESTS/logs/client19_pb25_20.log &
CLIENT_PID19=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_10.txt > $TESTS/logs/client1x_pb25_20.log &
CLIENT_PID20=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_1.txt > $TESTS/logs/client21_pb25_20.log &
CLIENT_PID21=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_2.txt > $TESTS/logs/client22_pb25_20.log &
CLIENT_PID22=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_3.txt > $TESTS/logs/client23_pb25_20.log &
CLIENT_PID23=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_4.txt > $TESTS/logs/client24_pb25_20.log &
CLIENT_PID24=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_5.txt > $TESTS/logs/client25_pb25_20.log &
CLIENT_PID25=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_6.txt > $TESTS/logs/client26_pb25_20.log &
CLIENT_PID26=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_7.txt > $TESTS/logs/client27_pb25_20.log &
CLIENT_PID27=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_8.txt > $TESTS/logs/client28_pb25_20.log &
CLIENT_PID28=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_9.txt > $TESTS/logs/client29_pb25_20.log &
CLIENT_PID29=$!
go run $CLIENT_GO -b=$TESTS/pure_basic_25_10.txt > $TESTS/logs/client2x_pb25_20.log &
CLIENT_PID20=$!

echo "[TEST] Waiting for client 1 to finish"
wait $CLIENT_PID01
echo "[TEST] Waiting for client 2 to finish"
wait $CLIENT_PID02
echo "[TEST] Waiting for client 3 to finish"
wait $CLIENT_PID03
echo "[TEST] Waiting for client 4 to finish"
wait $CLIENT_PID04
echo "[TEST] Waiting for client 5 to finish"
wait $CLIENT_PID05
echo "[TEST] Waiting for client 6 to finish"
wait $CLIENT_PID06
echo "[TEST] Waiting for client 7 to finish"
wait $CLIENT_PID07
echo "[TEST] Waiting for client 8 to finish"
wait $CLIENT_PID08
echo "[TEST] Waiting for client 9 to finish"
wait $CLIENT_PID09
echo "[TEST] Waiting for client 10 to finish"
wait $CLIENT_PID10

echo "[TEST] Waiting for client 1 to finish"
wait $CLIENT_PID11
echo "[TEST] Waiting for client 2 to finish"
wait $CLIENT_PID12
echo "[TEST] Waiting for client 3 to finish"
wait $CLIENT_PID13
echo "[TEST] Waiting for client 4 to finish"
wait $CLIENT_PID14
echo "[TEST] Waiting for client 5 to finish"
wait $CLIENT_PID15
echo "[TEST] Waiting for client 6 to finish"
wait $CLIENT_PID16
echo "[TEST] Waiting for client 7 to finish"
wait $CLIENT_PID17
echo "[TEST] Waiting for client 8 to finish"
wait $CLIENT_PID18
echo "[TEST] Waiting for client 9 to finish"
wait $CLIENT_PID19
echo "[TEST] Waiting for client 10 to finish"
wait $CLIENT_PID20

T="$(($(date +%s%N)-$T))"
M="$((T/1000000))"
echo "[TEST] Time Elapsed in milliseconds: ${M}"
echo "[TEST] Correctness not yet checked"
echo "[TEST] Finished."

