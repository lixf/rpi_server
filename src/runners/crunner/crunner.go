// A simple program which runs a client

package main

import (
    "flag"
    "fmt"
    "os"
    "client"
    "rpc/ipaddrs"
    "io/ioutil"
    "strings"
    "strconv"
    "errors"
)
//initialize flag variables
var test string
var time int
func init() {
    const (
        defaultTest = "default.txt"
        usageTest = "--benchmark  -b : give benchmark files"
        defaultTime = 0
        usageTime = "--time       -t : display timing information"
    )
    //for benchmark flags
    flag.StringVar(&test, "benchmark", defaultTest, usageTest)
    flag.StringVar(&test, "b", defaultTest, usageTest+ " (-b)")

    //for timing flags
    flag.IntVar(&time, "time", defaultTime, usageTime)
    flag.IntVar(&time, "t", defaultTime, usageTime+ " (-t)")
}

func checkError(err error) {
    if err != nil {
        fmt.Println("[CLIENT] [ERROR] ", err)
        os.Exit(1)
    }
}

func dispGet(client client.RpiClient, key string) {
    fmt.Println("[CLIENT] [GET] key: ", key)
    status, value, err := client.Get(key)
    checkError(err)
    fmt.Println("[CLIENT] [GET] Status: ", status, ", value: ", value)
}

func dispPost(client client.RpiClient, key, value string) {
    fmt.Println("[CLIENT] [POST] key: ", key, ", value:", value)
    status, err := client.Put(key, value)
    checkError(err)
    fmt.Println("[CLIENT] [POST] Status: ", status)
}

//have hashing capability
func dispCompute(client client.RpiClient, job,key,salt string,cost int) {
    fmt.Println("[CLIENT] [COMPUTE] job: ",job,"key ",key,"salt ",salt)
    status, res, err := client.Compute(job,key,salt,cost)
    checkError(err)
    fmt.Println("[CLIENT] [COMPUTE] Status: ", status, ", result: ", res)
}

func sendReq(client client.RpiClient, requests []string) error {
    for i:=0; i<len(requests)-1; i++ {
        req := requests[i]
        //split on space " "
        fields := strings.Split(req," ")
        cmd := fields[0]
        if strings.Contains(cmd,"COMPUTE") {
            //Parse the appropriate parameters of a compute job.
            jobType := fields[1]
            key := fields[2]
            salt := fields[3]
            cost,err := strconv.ParseInt(fields[4],10,0)
            checkError(err)
            //job,key,salt,cost
            dispCompute(client, jobType, key, salt, int(cost))
        } else if strings.Contains(cmd,"GET") {
            //key
            dispGet(client,fields[1])
        } else if strings.Contains(cmd,"POST") {
            //key,val
            dispPost(client,fields[1],fields[2])
        } else {
            fmt.Println(cmd)
            err := errors.New("Undefined request stream")
            return err
        }
    }
    return nil
}


func main() {
    flag.Parse()

    if flag.NArg() > 0 {
        flag.Usage()
        os.Exit(1)
    }

    fmt.Println("test file is ",test)
    fmt.Println("time? ",time)

    //parse the input file specified
    data, ioErr := ioutil.ReadFile(test)
    checkError(ioErr)

    fmt.Println("[CLIENT] parsing...")
    //do a split on newline to find requests
    requests := strings.Split(string(data),"\n")
    fmt.Println("[CLIENT] ... parsing finished.")

    //network code starts here
    fmt.Println("[CLIENT] Being created")
    client, err := client.NewClient(ipaddrs.MasterServerHostPort)
    checkError(err)
    fmt.Println("[CLIENT] Created successfully")

    //just parse and send the requests now
    fmt.Println("[CLIENT] Sending requests")
    sErr := sendReq(client,requests)
    checkError(sErr)

    //start sending messages from the file
    dispGet(client, "hihi")
    dispPost(client, "hihi", "asdf")
    dispGet(client, "hihi")
    dispPost(client, "hihi", "aaaaaaaaaaaaaaaaaaa")

    //TODO COMPUTE
    dispCompute(client, "hash","hihi","salt",1)
    dispGet(client, "hihi")
}

