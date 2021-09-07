package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func main() {
	
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		fmt.Printf("Message sent:%s\n",input)
		//On windows "\r\n" on other "\n". Extremely annoying. Use runtime.GOOS == "windows" to check. 
		if strings.TrimRight(input,"\r\n") == "quit" {
			return
		}
	}
}