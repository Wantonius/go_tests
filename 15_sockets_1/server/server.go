package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)
func main() {
	fmt.Println("Start server at port 5000")

	ln, _ := net.Listen("tcp", ":5000")
	fmt.Println("Waiting for connections!")
	conn, _ := ln.Accept()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from client:", string(message))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Message to client: ")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, text + "\n")
	
}