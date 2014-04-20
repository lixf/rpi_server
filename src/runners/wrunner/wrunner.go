package main

import (
    "fmt"
    "flag"
    "workerserver"
    "rpc/ipaddrs"
)

const defaultWorkerPort = ipaddrs.WorkerPort

var (
    port        = flag.String("port", defaultWorkerPort, "port number to listen on")
    numNodes    = flag.Int("N", 1, "the number of nodes in the ring")
    nodeID      = flag.Uint("id", 0, "a 32-bit unsigned node ID, used for consistent hashing")
)

func main() {
    flag.Parse()

    addr := ipaddrs.DetermineIP()
    _, err := workerserver.NewWorkerServer(addr + ":" + defaultWorkerPort)
    if err != nil {
        fmt.Println("[ERROR] Could not make worker server: ", err)
    }

    select {}
}
