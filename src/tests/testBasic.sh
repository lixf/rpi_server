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

go run $WORKER_GO -port=$WORKER_PORT1 > /dev/null &
go run $WORKER_GO -port=$WORKER_PORT2 > /dev/null &
go run $MASTER_GO -N=2



