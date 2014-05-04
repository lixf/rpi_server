package ipaddrs

import (
    "os"
    "net"
)

const MasterIP   = "192.168.1.4"
const MasterPort = "9001"
const MasterServerHostPort = MasterIP + ":" + MasterPort

const WorkerPort = "8000"
const NumNodes = 2

func DetermineIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        os.Stderr.WriteString("Oops: " + err.Error() + "\n")
        os.Exit(1)
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}
