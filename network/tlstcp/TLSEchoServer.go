package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1201"
	crt, err := tls.LoadX509KeyPair("myserver.crt", "myserver.key")
	checkError(err)
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	listener, err := tls.Listen("tcp", service, tlsConfig)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte
	for {
		// read up to 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		// write the n bytes read
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
