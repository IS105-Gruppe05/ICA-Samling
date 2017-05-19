package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")

	// start server
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		errors.New("Problem setting up server.")
	}

	// continuously accept connections
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			errors.New("Error connecting.")
		}
		defer conn.Close()
		go handleConnection(conn)
	}
}

func handleConnection(c *net.TCPConn) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	for {
		msg := make([]byte, 32)
		// read the sent message from the connection
		c.Read(msg)
		fmt.Printf("%x\n", msg)

		// echo it back to the client
		fmt.Fprintf(c, "Echoing back %x at %s", msg,
			time.Now().Format(layout))
	}
}
