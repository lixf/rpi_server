//This is the API for a client in the RPI server system

package client

import "rpc/masterrpc"

type RpiClient interface {
    Get(key string) (masterrpc.Status, string, error)
    Put(key, value string) (masterrpc.Status, error)
    //TODO COMPUTE: Feel free to change argument or result
    Compute(job string,key string,salt string,cost int) (masterrpc.Status, string, error)
}

