package workerserver

import (
    "rpc/workerrpc"
    "rpc/masterrpc"
    "rpc/ipaddrs"
    "net"
    "net/http"
    "net/rpc"
    "hashing"
    "time"
    "fmt"
    "sync"
)

type workerServer struct {
    storageMap  map[string]string
    itemLock    sync.Mutex
    myNode      *masterrpc.Node
}

const masterServerHostPort = ipaddrs.MasterServerHostPort

func NewWorkerServer(workerServerHostPort string) (WorkerServer, error) {
    fmt.Println("[WORKER] New worker")
    ws := new(workerServer)
    ws.storageMap = make(map[string]string)
    ws.myNode = &masterrpc.Node{}
    ws.myNode.NodeID = hashing.HashString(workerServerHostPort)
    ws.myNode.HostPort = workerServerHostPort

    listener, err := net.Listen("tcp", workerServerHostPort)
    if err != nil {
        return nil, err
    }
    err = rpc.RegisterName("WorkerServer", workerrpc.Wrap(ws))
    if err != nil {
        return nil, err
    }

    //Achieve a connection with the master node
    master, err := rpc.DialHTTP("tcp", masterServerHostPort)
    for err != nil {
        time.Sleep(1 * time.Second)
        fmt.Println("[WORKER] Cannot see master")
        master, err = rpc.DialHTTP("tcp", masterServerHostPort)
    }
    fmt.Println("[WORKER] Worker has successfully connected to master")
    rpc.HandleHTTP()
    go http.Serve(listener, nil)
    //Register self with master node
    args := &masterrpc.RegisterArgs{ServerInfo: *ws.myNode}
    var reply masterrpc.RegisterReply
    err = master.Call("MasterServer.RegisterServer", args, &reply)
    for (err != nil) {
        time.Sleep(1 * time.Second)
        fmt.Println("[WORKER] Registering with master")
        err = master.Call("MasterServer.RegisterServer", args, &reply)
    }
    ws.myNode.NodeID = reply.ServerInfo.NodeID //Assigned ID
    return ws, nil
}

func (ws *workerServer) Get(args *workerrpc.GetArgs, reply *workerrpc.GetReply) error {
    fmt.Println("[WORKER] GET called")
    //TODO change itemlock to R/W lock
    ws.itemLock.Lock()
    val, present := ws.storageMap[args.Key]
    if present {
        reply.Status = workerrpc.OK
        reply.Value = val
    } else {
        reply.Status = workerrpc.ItemNotFound
    }
    ws.itemLock.Unlock()
    return nil
}

func (ws *workerServer) Put(args *workerrpc.PutArgs, reply *workerrpc.PutReply) error {
    fmt.Println("[WORKER] PUT called")
    ws.itemLock.Lock()

    ws.storageMap[args.Key] = args.Value

    reply.Status = workerrpc.OK
    ws.itemLock.Unlock()
    return nil
}

//TODO COMPUTE
func (ws *workerServer) Compute(args *workerrpc.ComputeArgs, reply *workerrpc.ComputeReply) error {
    fmt.Println("[WORKER] COMPUTE called")

    reply.Result = args.Param + ", NOW WITH MORE COMPUTATION!"

    reply.Status = workerrpc.OK
    return nil
}


