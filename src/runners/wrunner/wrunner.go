package main

import (
    "fmt"
    "flag"
    "net"
    "os"
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

    name, err := os.Hostname()
    if err != nil {
        fmt.Println("[ERROR]", err)
    } else {
        fmt.Println("Worker name: ", name)
    }
    addrs, err := net.LookupHost(name)
    if err != nil {
        fmt.Println("[ERROR]", err)
    } else {
        fmt.Println("Address: ", addrs[0])
    }

    _, err = workerserver.NewWorkerServer(addrs[0] + ":" + defaultWorkerPort)
    if err != nil {
        fmt.Println("[ERROR] Could not make worker server: ", err)
    }

    select {}
}
