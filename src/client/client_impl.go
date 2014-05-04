package client

import (
    "fmt"
    "time"
    "net/rpc"
    "rpc/masterrpc"
    "io/ioutil"
)

type rpiClient struct {
    client *rpc.Client
}

func NewClient(masterServerHostPort string) (RpiClient, error) {
    fmt.Println("[CLIENT], ", time.Now())
    cli, err := rpc.DialHTTP("tcp", masterServerHostPort)
    if err != nil {
        return nil, err
    }
    return &rpiClient{client: cli}, nil
}

func (rc *rpiClient) Get(key string) (masterrpc.Status, string, error) {
    args := &masterrpc.GetArgs{Key: key}
    var reply masterrpc.GetReply
    if err := rc.client.Call("MasterServer.Get", args, &reply); err != nil {
        return 0, "", err
    }
    return reply.Status, reply.Value, nil
}

func (rc *rpiClient) Put(key, value string) (masterrpc.Status, error) {
    args := &masterrpc.PutArgs{Key: key, Value: value}
    var reply masterrpc.PutReply
    if err := rc.client.Call("MasterServer.Put", args, &reply); err != nil {
        return 0, err
    }
    return reply.Status, nil
}

func (rc *rpiClient) Hash(key string, salt string, cost int) (masterrpc.Status, string, error){
    args := &masterrpc.HashArgs{Key: key, Salt: salt, Cost: cost}
    var reply masterrpc.HashReply
    //TODO COMPUTE: Pass real arguments, get a real response
    if err := rc.client.Call("MasterServer.Hash", args, &reply); err != nil {
        return 0, "", err
    }
    return reply.Status, reply.Result, nil
}

func (rc *rpiClient) Pict(local string, store string) (masterrpc.Status, string, error){
    path := "src/local_pict/"
    //do file io here and read in the picture stored at "local"
    pbytes, ferr := ioutil.ReadFile(path+local);
    if ferr != nil {
        fmt.Println("file opening err")
        return 0,"",ferr
    }
    //pass the picture over the network
    args := &masterrpc.PictArgs{PictBytes: pbytes, Store: store}
    var reply masterrpc.PictReply
    //COMPUTE: Pass real arguments, get a real response
    if err := rc.client.Call("MasterServer.Pict", args, &reply); err != nil {
        return 0, "", err
    }
    return reply.Status, reply.Result, nil
}
