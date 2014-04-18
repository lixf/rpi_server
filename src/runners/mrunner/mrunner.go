package main

import (
    "fmt"
    "flag"
    "rpc/ipaddrs"
    "masterserver"
)

const defaultMasterPort = ipaddrs.MasterPort

var (
    port        = flag.String("port", defaultMasterPort, "port number to listen on")
    numNodes    = flag.Int("N", ipaddrs.NumNodes, "the number of nodes in the ring")
    nodeID      = flag.Uint("id", 0, "a 32-bit unsigned node ID, used for consistent hashing")
)

func main() {
    flag.Parse()

    _, err := masterserver.NewMasterServer(*port, *numNodes)
    if err != nil {
        fmt.Println("[ERROR] Could not make master server")
    }

    select {}
}
