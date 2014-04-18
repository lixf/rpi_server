//This is the API for a client in the RPI server system

package client

import "rpc/masterrpc"

type RpiClient interface {
    //TODO probably include a GET/PUT/COMPUTE
    Get(key string) (masterrpc.Status, string, error)
    Put(key, value string) (masterrpc.Status, error)
}

