package masterserver

import (
    "fmt"
    "net"
    "net/http"
    "net/rpc"
    "rpc/masterrpc"
    "rpc/workerrpc"
    "hashing"
    "math"
    "errors"
)

type masterServer struct {
    //TODO Fill in
    serverReady     bool
    numNodes        int
    numReadyNodes   int
    serversReadyMap map[uint32]bool
    servers         []masterrpc.Node //For viewing info about workers
    connections     []*rpc.Client
}

/* Function which converts index into list of workers (called "servers") 
 * to an appropriate hash value. 
 */
func indexToHash(ms *masterServer, index uint32) uint32 {
    return uint32(float32(index)*float32(math.MaxUint32)/float32(ms.numNodes))
}

/* Function which converts a hash value into an appropriate index into 
 * the list of workers. This is used to determine which worker handles
 * a particular hash value
 */
func (ms *masterServer) hashToIndex(hash uint32) int {
    //XXX Note, this is linear, but can debug. TODO: SHOULD NOT ITERATE
    for index, node := range ms.servers {
        lowerID := node.NodeID
        var upperID uint32
        if index == len(ms.servers) - 1 {
            upperID = math.MaxUint32
        } else {
            upperID = ms.servers[index + 1].NodeID
        }
        if lowerID <= hash && upperID < hash {
            //Inclusive --  [index, nextIndex) -- Exclusive
            return index
        } else if (upperID == math.MaxUint32) && (upperID == hash) {
            //(exception -- maxInt is included in "nextIndex")
            return index
        }
    }
    //XXX Should assert false
    return 0
}

func NewMasterServer(masterServerHostPort string, numNodes int) (MasterServer, error) {
    fmt.Println("[MASTER] New master server")
    ms := new(masterServer)
    ms.serverReady = false
    ms.numNodes = numNodes
    ms.numReadyNodes = 0
    ms.serversReadyMap = make(map[uint32]bool)
    ms.servers = make([]masterrpc.Node, 0)
    ms.connections = make([]*rpc.Client, 0)

    listener, err := net.Listen("tcp", ":" + masterServerHostPort)
    if err != nil {
        fmt.Println("[MASTER][ERROR]", err)
        return nil, err
    }

    err = rpc.RegisterName("MasterServer", masterrpc.Wrap(ms))
    if err != nil {
        fmt.Println("[MASTER][ERROR]", err)
        return nil, err
    }

    rpc.HandleHTTP()
    go http.Serve(listener, nil)
    fmt.Println("[MASTER] Booted successfully -- waiting for workers")
    return ms, nil
}

func (ms *masterServer) RegisterServer(args *masterrpc.RegisterArgs, reply *masterrpc.RegisterReply) error {
    id := args.ServerInfo.NodeID
    fmt.Println("[MASTER] Register Server called")
    if _, present := ms.serversReadyMap[id]; !present {
        smallID := uint32(len(ms.serversReadyMap))
        hashID := indexToHash(ms, smallID)
        args.ServerInfo.NodeID = hashID
        fmt.Println("[MASTER] assigning ID ", args.ServerInfo.NodeID)
        fmt.Println("[MASTER] Calling: ", args.ServerInfo.HostPort)
        cli, err := rpc.DialHTTP("tcp", args.ServerInfo.HostPort)
        if err != nil {
            fmt.Println("[ERROR] ", err)
            return errors.New("Cannot dial worker")
        }
        fmt.Println("[MASTER] Registered with worker")
        ms.connections = append(ms.connections, cli)
        ms.servers = append(ms.servers, args.ServerInfo)
        ms.serversReadyMap[id] = true
    }
    if len(ms.serversReadyMap) == ms.numNodes {
        reply.Status = masterrpc.OK
        if !ms.serverReady {
            //TODO setup
            fmt.Println("[MASTER] Ready: ", ms.servers)
            ms.serverReady = true
        }
    } else {
        fmt.Println("Waiting for ", ms.numNodes - len(ms.serversReadyMap), "more nodes")
        reply.Status = masterrpc.NotReady
    }
    reply.ServerInfo = args.ServerInfo
    return nil
}

func (ms *masterServer) Get(args *masterrpc.GetArgs, reply *masterrpc.GetReply) error {
    fmt.Println("[MASTER] GET")
    if !ms.serverReady {
        reply.Status = masterrpc.NotReady
    }
    //TODO Find the appropriate worker, get. Possibly cache later.
    index := ms.hashToIndex(hashing.HashString(args.Key))
    cli := ms.connections[index]
    wArgs := &workerrpc.GetArgs{Key: args.Key}
    var wReply workerrpc.GetReply
    if err := cli.Call("WorkerServer.Get", wArgs, &wReply); err != nil {
        return err
    } else {
        reply.Value = wReply.Value
        return nil
    }
}

func (ms *masterServer) Put(args *masterrpc.PutArgs, reply *masterrpc.PutReply) error {
    fmt.Println("[MASTER] PUT")
    if !ms.serverReady {
        reply.Status = masterrpc.NotReady
    }
    //TODO Find the appropriate worker, put. Possibly cache later.
    index := ms.hashToIndex(hashing.HashString(args.Key))
    cli := ms.connections[index]
    wArgs := &workerrpc.PutArgs{Key: args.Key, Value: args.Value}
    var wReply workerrpc.PutReply
    if err := cli.Call("WorkerServer.Put", wArgs, &wReply); err != nil {
        return err
    } else {
        return nil
    }
}
