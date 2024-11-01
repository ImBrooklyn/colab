package network

import (
    "fmt"
    "net"
)

func InitNetwork(port int) net.Listener {
    host := "0.0.0.0"
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
    if err != nil {
        panic("failed to listen:" + err.Error())
    }
    return listener
}
