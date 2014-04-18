// A simple program which runs a client

package main

import (
    "flag"
    "fmt"
    "os"
    "client"
    "rpc/ipaddrs"
)


func init() {
    flag.Usage = func() {
        fmt.Fprintln(os.Stderr, "The crunner program creates a client.")
        fmt.Fprintln(os.Stderr, "You can control it with commands: ")
        fmt.Fprintln(os.Stderr, "UNIMPLEMENTED")
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Println("[CLIENT] [ERROR] ", err)
        os.Exit(1)
    }
}

func dispGet(client client.RpiClient, key string) {
    fmt.Println("[CLIENT] [GET] key: ", key)
    status, value, err := client.Get(key)
    checkError(err)
    fmt.Println("[CLIENT] [GET] Status: ", status, ", value: ", value)
}

func dispPost(client client.RpiClient, key, value string) {
    fmt.Println("[CLIENT] [POST] key: ", key, ", value:", value)
    status, err := client.Put(key, value)
    checkError(err)
    fmt.Println("[CLIENT] [POST] Status: ", status)
}

func main() {
    /* 
    TODO Figure out how to take command line args
    flag.Parse()
    if flag.NArg() < 2 {
        flag.Usage()
        os.Exit(1)
    }
    */

    client, err := client.NewClient(ipaddrs.MasterServerHostPort)
    checkError(err)

    dispGet(client, "hihi")
    dispPost(client, "hihi", "asdf")
    dispGet(client, "hihi")
    dispPost(client, "hihi", "aaaaaaaaaaaaaaaaaaa")
}


