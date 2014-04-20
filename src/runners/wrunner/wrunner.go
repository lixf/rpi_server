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
)

func main() {
    flag.Parse()

    addr := ipaddrs.DetermineIP()
    _, err := workerserver.NewWorkerServer(addr + ":" + *port)
    if err != nil {
        fmt.Println("[ERROR] Could not make worker server: ", err)
    }

    select {}
}
