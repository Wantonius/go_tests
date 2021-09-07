package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {

	connection, _ := net.Dial("tcp", "127.0.0.1:5000")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Message to server: ")
    text, _ := reader.ReadString('\n')
    fmt.Fprintf(connection, text + "\n")
    message, _ := bufio.NewReader(connection).ReadString('\n')
    fmt.Print("Message from server: "+message)
  
}