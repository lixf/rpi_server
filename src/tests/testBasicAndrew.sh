#!/bin/bash

echo "Basic Test: Two workers, one master, one client."

USERNAME=smklein
MASTER_IP=unix2.andrew.cmu.edu
WORKER_IP1=unix5.andrew.cmu.edu
WORKER_IP2=unix6.andrew.cmu.edu


PROJECT_PATH=/afs/andrew.cmu.edu/usr23/smklein/private/15418/rpi_server
#TODO along with the master port, you may need to change this. I'm bad at closing servers.
WORKER_PORT1=8050
WORKER_PORT2=8060
WORKER_GO=$PROJECT_PATH/src/runners/wrunner/wrunner.go
MASTER_GO=$PROJECT_PATH/src/runners/mrunner/mrunner.go
LOGS=/afs/andrew.cmu.edu/usr23/smklein/private/15418/rpi_server/src/tests/logs

echo "WORKER 1:"
ssh $USERNAME@$WORKER_IP1 "go run $WORKER_GO -port=$WORKER_PORT1 > $LOGS/worker1.log &" < /dev/null 
echo "WORKER 2:"
ssh -T $USERNAME@$WORKER_IP2 "go run $WORKER_GO -port=$WORKER_PORT2 > $LOGS/worker2.log &" < /dev/null
echo "MASTER:" 
ssh -T $USERNAME@$MASTER_IP "go run $MASTER_GO -N=2 > $LOGS/master.log &" < /dev/null
echo "CLIENT (run locally):"
CLIENT_GO=$GOPATH/src/runners/crunner/crunner.go
TESTS=$GOPATH/src/tests
go run $CLIENT_GO -b=$TESTS/basicGetPost.txt > $TESTS/logs/client.log &

echo "Waiting for tests to execute"
sleep 5
echo "Waited for five seconds. Fetching results"

scp $USERNAME@$WORKER_IP1:$LOGS/worker1.log logs/worker1.log
scp $USERNAME@$WORKER_IP2:$LOGS/worker2.log logs/worker2.log
scp $USERNAME@$MASTER_IP:$LOGS/master.log logs/master.log

