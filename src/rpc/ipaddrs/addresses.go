package ipaddrs

import (
    "os"
    "net"
)
//Unix2
//Pi master
const MasterIP   = "192.168.1.51"
//Sean's Laptop
//const MasterIP   = "192.168.1.51"
//Unix
//const MasterIP = "128.2.13.133"
const MasterPort = "9003"
const MasterServerHostPort = MasterIP + ":" + MasterPort

//Used only by workers, as a default
const WorkerPort = "8000"
const NumNodes = 3

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
