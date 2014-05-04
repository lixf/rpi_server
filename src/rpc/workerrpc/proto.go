package workerrpc

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
    Servers []Node
}

type GetServersArgs struct {
    //Empty
}

type GetServersReply struct {
    Status  Status
    Servers []Node
}

type GetArgs struct {
    Key       string
    HostPort  string
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

//TODO COMPUTE
//added hasing with salt
type HashArgs struct{
    Salt  string    //Effectively random
    Cost  int       //Number of hashing iterations
    Key   string    //... yeah.
}
type HashReply struct {
    Status Status
    Result string
}

//transmitting pictures
type PictArgs struct{
    PictBytes []byte
    Store     string

}
type PictReply struct {
    Status Status
    Result string
}
