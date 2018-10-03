package main

import (
        "net"
        "log"
        "strings"
        "encoding/binary"
        "github.com/hekmon/transmissionrpc"
)

func checkErr(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
        ifi, err := net.InterfaceByName("tun0")
        checkErr(err)
        addrs, err := ifi.Addrs()
        checkErr(err)

        for _, addr := range addrs {
                // Check for IPv4
                if strings.Contains(addr.String(), ".") {
                        ipv4Addr, _, err := net.ParseCIDR(addr.String())
                        checkErr(err)
                        port := 30000 + int64(binary.BigEndian.Uint64(ipv4Addr[8:]) & 0xFFF)
                        transmission := transmissionrpc.New("127.0.0.1", "transmission", "transmission", nil)
                        sessionArgs, err := transmission.SessionArgumentsGet()
                        checkErr(err)
                        if *sessionArgs.PeerPort != port {
                                transmission.SessionArgumentsSet(&transmissionrpc.SessionArguments{PeerPort: &port})
                        }
                }
        }
}
