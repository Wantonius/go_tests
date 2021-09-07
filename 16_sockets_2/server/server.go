package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func handleClient(conn net.Conn) {
	
	var stringbuffer string
	for stringbuffer != "quit" {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')	
		if err != nil {
			fmt.Println("Client left.")
			conn.Close()
			return
		}
		stringbuffer = string(buffer[:len(buffer)-1])
		fmt.Println("Client message:", stringbuffer)
	}

}

func main() {
	fmt.Println("Accepting connections at 5000")
	l, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()


	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleClient(c)
	}
}
