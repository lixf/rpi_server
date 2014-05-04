package workerserver

import "rpc/workerrpc"

type WorkerServer interface {
    Get(*workerrpc.GetArgs, *workerrpc.GetReply) error
    Put(*workerrpc.PutArgs, *workerrpc.PutReply) error
    //COMPUTE
    Hash(*workerrpc.HashArgs, *workerrpc.HashReply) error
    Pict(*workerrpc.PictArgs, *workerrpc.PictReply) error
}
