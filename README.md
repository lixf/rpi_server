rpi_server
==========

418 final project - A Distributed, Low-Power Server, meant to be run on raspberry pis.

By Xiaofan and Sean!

From the src directory...
    To acivate a master:
        > go run runners/mrunner/mrunner.go
    To acivate a worker:
        > go run runners/wrunner/wrunner.go
    To acivate a client:
        > go run runners/crunner/crunner.go


src/client:
    Client interface to access server.
src/masterserver:
    Main master server. Used to reroute client requests to many 
    workers -- may use caching of results in the future. Only point of contact with 
    with workers.
src/workerserver:
    Worker servers -- store information, do computation.
src/rpc/masterrpc:
src/rpc/workerrpc:
    Contains information regarding rpc calls to master (from client + workers) and 
    rpc calls to worker (from master).

!!!
src/rpc/ipaddrs:
    Contains information about port choices, IP address, and number of workers.
    Must be customized for configuration, dependent on master IP address.
!!!

src/hashing:
    Contains information regarding hashing functionality.
src/runners:
    Contains code to actually activate client, master, worker.

