#!/bin/bash

echo "Basic Test: Three workers, one master, three clients."

USERNAME=pi
MASTER_IP=192.168.1.51
WORKER_IP1=192.168.1.52
WORKER_IP2=192.168.1.53
WORKER_IP3=192.168.1.54
PROJECT_PATH=~/code/rpi_server

#TODO along with the master port, you may need to change this. I'm bad at closing servers.
WORKER_PORT1=$(((RANDOM % 10000) + 10000))
WORKER_PORT2=$(((RANDOM % 10000) + 10000))
WORKER_PORT3=$(((RANDOM % 10000) + 10000))
WORKER_GO=$PROJECT_PATH/src/runners/wrunner/wrunner.go
MASTER_GO=$PROJECT_PATH/src/runners/mrunner/mrunner.go
LOGS=$PROJECT_PATH/src/tests/logs

echo "[TEST] MASTER:" 
ssh -T $USERNAME@$MASTER_IP "go run $MASTER_GO -N=2 > $LOGS/master.log &" < /dev/null
sleep 3
echo "[TEST] WORKER 1:"
ssh $USERNAME@$WORKER_IP1 "go run $WORKER_GO -port=$WORKER_PORT1 > $LOGS/worker1.log &" < /dev/null 
sleep 3
echo "[TEST] WORKER 2:"
ssh -T $USERNAME@$WORKER_IP2 "go run $WORKER_GO -port=$WORKER_PORT2 > $LOGS/worker2.log &" < /dev/null
sleep 3
echo "[TEST] WORKER 3:"
ssh -T $USERNAME@$WORKER_IP3 "go run $WORKER_GO -port=$WORKER_PORT3 > $LOGS/worker3.log &" < /dev/null

echo "[TEST] Waiting for master/worker system to set up..."
sleep 3
echo "[TEST] Starting timer:"
T="$(date +%s%N)"

echo "[TEST] CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests

go run $CLIENT_GO -b=$TESTS/shortGetPost1.txt > $TESTS/logs/client1.log &
CLIENT_PID1=$!
go run $CLIENT_GO -b=$TESTS/shortGetPost2.txt > $TESTS/logs/client2.log &
CLIENT_PID2=$!
go run $CLIENT_GO -b=$TESTS/shortGetPost3.txt > $TESTS/logs/client3.log &
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

scp $USERNAME@$WORKER_IP1:$LOGS/worker1.log logs/worker1.log
scp $USERNAME@$WORKER_IP2:$LOGS/worker2.log logs/worker2.log
scp $USERNAME@$WORKER_IP3:$LOGS/worker3.log logs/worker3.log
scp $USERNAME@$MASTER_IP:$LOGS/master.log logs/master.log

