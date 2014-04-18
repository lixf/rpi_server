package workerserver

import "rpc/workerrpc"

type WorkerServer interface {
    Get(*workerrpc.GetArgs, *workerrpc.GetReply) error

    Put(*workerrpc.PutArgs, *workerrpc.PutReply) error
}
