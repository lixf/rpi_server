package client

import (
    "fmt"
    "time"
    "net/rpc"
    "rpc/masterrpc"
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

func (rc *rpiClient) Compute(job string, key string, salt string, cost int) (masterrpc.Status, string, error){
    args := &masterrpc.ComputeArgs{Job: job, Key: key, Salt: salt, Cost: cost}
    var reply masterrpc.ComputeReply
    //TODO COMPUTE: Pass real arguments, get a real response
    if err := rc.client.Call("MasterServer.Compute", args, &reply); err != nil {
        return 0, "", err
    }
    return reply.Status, reply.Result, nil
}

