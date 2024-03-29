package main

import (
    "fmt"
    "flag"
    "os"
    "rpc/ipaddrs"
    "masterserver"
)

const defaultMasterPort = ipaddrs.MasterPort

var (
    port        = flag.String("port", defaultMasterPort, "port number to listen on")
    numNodes    = flag.Int("N", ipaddrs.NumNodes, "the number of nodes in the ring")
)

func main() {
    flag.Parse()

    _, err := masterserver.NewMasterServer(*port, *numNodes)
    if err != nil {
        fmt.Println("[ERROR] Could not make master server")
        os.Exit(1)
    }
    select {}
}
