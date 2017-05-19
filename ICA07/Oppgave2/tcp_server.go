package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
