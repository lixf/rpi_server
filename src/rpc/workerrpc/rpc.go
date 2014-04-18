package workerrpc

type RemoteWorkerServer interface {
    Get(*GetArgs, *GetReply) error
    Put(*PutArgs, *PutReply) error
}

type WorkerServer struct {
    RemoteWorkerServer
}

func Wrap(w RemoteWorkerServer) RemoteWorkerServer {
    return &WorkerServer{w}
}
