package masterrpc

type RemoteMasterServer interface {
    RegisterServer(*RegisterArgs, *RegisterReply) error
    Get(*GetArgs, *GetReply) error
    Put(*PutArgs, *PutReply) error
    //TODO COMPUTE
    Hash(*HashArgs, *HashReply) error
    Pict(*PictArgs, *PictReply) error
}

type MasterServer struct {
    RemoteMasterServer
}

func Wrap(w RemoteMasterServer) RemoteMasterServer {
    return &MasterServer{w}
}
