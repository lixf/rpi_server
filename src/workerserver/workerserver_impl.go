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
    "errors"
    "lib/code.google.com/p/go.crypto/bcrypt"//for hashing
    "io/ioutil" //for writing files
)

type workerServer struct {
    storageMap  map[string]string
    itemLock    sync.Mutex
    myNode      *masterrpc.Node
}

const masterServerHostPort = ipaddrs.MasterServerHostPort

func NewWorkerServer(workerServerHostPort string) (WorkerServer, error) {
    fmt.Println("[WORKER] New worker", time.Now())
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
        fmt.Println("[WORKER] Cannot see master", err)
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



///////////COMPUTE//////////////
//functionalities:
//1. hasing with salt
func (ws *workerServer) salt_hash (key string, salt string, cost int) error{
    fmt.Println("[WORKER] COMPUTE->hashing")
    //take the lock on the storage
    //we could also make a new table to prevent messing up 
    //hased with unhashed values
    ws.itemLock.Lock()
    val, present := ws.storageMap[key]
    if present {
        //use this val to do computation uses bcrypt
        salted := []byte(val + salt)
        //this function takes n as the cost for the hashing
        h, err := bcrypt.GenerateFromPassword(salted,cost)
        //check error and then put the hash back
        if err != nil {
          ws.itemLock.Unlock()
          return err
        }
        fmt.Println("h is ",h)
        ws.storageMap[key] = string(h)
    } else {
      ws.itemLock.Unlock()
      err := errors.New("can't find with key")
      return err
    }
    //unlock here
    ws.itemLock.Unlock()
    return nil
}


//2. GPU image rendering


//3. compute prime




///Wrappers for each compute here
///1.HASH -- calls salt_hash
func (ws *workerServer) Hash(args *workerrpc.HashArgs, reply *workerrpc.HashReply) error {
    fmt.Println("[WORKER] HASH called")
    //simulating a password hashing with salt
    cost := args.Cost
    salt := args.Salt

    //looking up value same as in GET
    hash_err := ws.salt_hash (args.Key,salt,cost)
    if hash_err != nil {
      return hash_err
    }
    reply.Result = "success"
    reply.Status = workerrpc.OK
    return nil
}


///2.Pict
func (ws *workerServer) Pict(args *workerrpc.PictArgs, reply *workerrpc.PictReply) error {
    fmt.Println("[WORKER] PICT called")
    //get the arguments and write to file
    pbytes := args.PictBytes
    store  := args.Store
    path   := "src/store_pict/"

    //writes to file filemode = rwrwrw
    ferr := ioutil.WriteFile((path+store), pbytes, 0666)

    if ferr != nil {
      return ferr
    }
    reply.Result = "success"
    reply.Status = workerrpc.OK
    return nil
}
