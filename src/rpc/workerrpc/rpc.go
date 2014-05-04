package workerrpc

type RemoteWorkerServer interface {
    Get(*GetArgs, *GetReply) error
    Put(*PutArgs, *PutReply) error
    //TODO COMPUTE
    Hash(*HashArgs, *HashReply) error
    Pict(*PictArgs, *PictReply) error
}

type WorkerServer struct {
    RemoteWorkerServer
}

func Wrap(w RemoteWorkerServer) RemoteWorkerServer {
    return &WorkerServer{w}
}
