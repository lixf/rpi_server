package main

import (
    "encoding/json"
    "fmt"
    "net"
    "os"
)

func handleConnection(conn net.Conn){
    mesg := "INCOMING MESSAGE -- Hello from server"
    b, err := json.Marshal(mesg)
    if err != nil {
        fmt.Println("Cannot marshal outgoing message")
    }
    conn.Write(b)
    fmt.Println("Server has written message")
}

func main() {
    fmt.Println("New Server Starting")
    //The listen function creates servers
    ln, err := net.Listen("tcp", "128.2.13.145:8080")
    if err != nil {
        fmt.Println("Server unable to listen")
        os.Exit(-1)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        } else {
            go handleConnection(conn)
        }
    }
}
