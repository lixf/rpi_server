#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client01_p25_20.log &
CLIENT_PID01=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client02_p25_20.log &
CLIENT_PID02=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client03_p25_20.log &
CLIENT_PID03=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client04_p25_20.log &
CLIENT_PID04=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client05_p25_20.log &
CLIENT_PID05=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client06_p25_20.log &
CLIENT_PID06=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client07_p25_20.log &
CLIENT_PID07=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client08_p25_20.log &
CLIENT_PID08=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client09_p25_20.log &
CLIENT_PID09=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client0x_p25_20.log &
CLIENT_PID10=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client11_p25_20.log &
CLIENT_PID11=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client12_p25_20.log &
CLIENT_PID12=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client13_p25_20.log &
CLIENT_PID13=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client14_p25_20.log &
CLIENT_PID14=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client15_p25_20.log &
CLIENT_PID15=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client16_p25_20.log &
CLIENT_PID16=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client17_p25_20.log &
CLIENT_PID17=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client18_p25_20.log &
CLIENT_PID18=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client19_p25_20.log &
CLIENT_PID19=$!
go run $CLIENT_GO -b=$TESTS/pict_25.txt > $TESTS/logs/client1x_p25_20.log &
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

