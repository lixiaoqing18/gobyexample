package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name) //default ipv4,"ip6" for ipv6
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}

	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("CNAME: ", cname)

	tcpAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("TcpAddr: ", tcpAddr)
	os.Exit(0)
}
