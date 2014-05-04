package masterrpc

type Status int

const (
    OK      Status = iota + 1
    NotReady
    WrongServer
    ItemNotFound
)

type Node struct {
    HostPort string
    NodeID   uint32
}

type RegisterArgs struct {
    ServerInfo Node
}

type RegisterReply struct {
    Status  Status
    ServerInfo Node
}

type GetServersArg struct {
    //Empty
}

type GetServersReply struct {
    Status  Status
    Servers []Node
}

type GetArgs struct {
    Key       string
}

type GetReply struct {
    Status Status
    Value string
}

type PutArgs struct {
    Key string
    Value string
}

type PutReply struct {
    Status Status
}

// COMPUTE
//added hasing with salt
//compute test
type HashArgs struct{
    Salt  string
    Cost  int
    Key   string
}
type HashReply struct {
    Status Status
    Result string
}

//added pictrue transmitting
//bandwidth test
type PictArgs struct{
    PictBytes   []byte
    Store       string
}
type PictReply struct {
    Status Status
    Result string
}
