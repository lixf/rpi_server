package client

import (
    "fmt"
    "testing"
    "masterserver"
    "workerserver"
)

type testFramework struct {
    t       *testing.T
    master  masterserver.MasterServer
    workers []workerserver.WorkerServer
    clients []RpiClient
    numClients  int
    numWorkers  int


}

func (tf testFramework) chk(err error) {
    if err != nil {
        tf.t.Error(err)
    }
}

func newTestFramework(t *testing.T, numClients, numWorkers int) *testFramework{
    tf := new(testFramework)
    tf.t = t
    
    portNum := "8000"
    ms, err := masterserver.NewMasterServer(portNum, numWorkers)
    tf.chk(err)
    
    tf.master = ms
    tf.workers = nil//XXX
    tf.clients = nil//XXX
    tf.numClients = numClients
    tf.numWorkers = numWorkers

    return tf
}


func TestBasic1(t *testing.T){
    fmt.Println("TestBasic 1")
    tf := newTestFramework(t, 2, 2)
    if tf == nil {
        fmt.Println("WTF")
    }
}
