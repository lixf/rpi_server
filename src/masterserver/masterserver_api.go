package masterserver

import "rpc/masterrpc"

type MasterServer interface {
    RegisterServer(*masterrpc.RegisterArgs, *masterrpc.RegisterReply) error
    Get(*masterrpc.GetArgs, *masterrpc.GetReply) error
    Put(*masterrpc.PutArgs, *masterrpc.PutReply) error
    //TODO COMPUTE
    Hash(*masterrpc.HashArgs, *masterrpc.HashReply) error
}


