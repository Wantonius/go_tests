package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

type Manager struct {
    clients    	map[*Client]bool
    broadcast  	chan string
    register   	chan *Client
    unregister 	chan *Client
}

type Client struct {
    socket 	net.Conn
    data   	chan string
	name  	string
}

func (manager *Manager) startService() {
    for {
        select {
        case connection := <-manager.register:
            manager.clients[connection] = true
            fmt.Printf("User %s has entered the chat!\n",connection.name)
        case connection := <-manager.unregister:
            if _, ok := manager.clients[connection]; ok {
                close(connection.data)
				fmt.Printf("User %s has left the chat!\n",connection.name)
                delete(manager.clients, connection)
               
            }
        case message := <-manager.broadcast:
            for connection := range manager.clients {
                select {
                case connection.data <- message:
                default:
                    close(connection.data)
                    delete(manager.clients, connection)
                }
            }
        }
    }
}

func (manager *Manager) receive(client *Client) {
    for {
        message := make([]byte, 4096)
        length, err := client.socket.Read(message)
        if err != nil {
            manager.unregister <- client
            client.socket.Close()
            break
        }
        if length > 0 {
			temp_message := client.name+" "+string(message)
			final_message := strings.TrimRight(temp_message,"\n")
            manager.broadcast <- final_message
        }
    }
}

func (manager *Manager) send(client *Client) {
    defer client.socket.Close()
    for {
        select {
        case message, ok := <-client.data:
            if !ok {
                return
            }
            fmt.Fprintf(client.socket,message)
        }
    }
}

func main() {
    fmt.Println("Starting server at port 5000")
    listener, error := net.Listen("tcp", "localhost:5000")
    if error != nil {
        fmt.Println(error)
    }
    manager := Manager{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan string),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
    go manager.startService()
    for {
		fmt.Println("Waiting for new clients")
        connection, _ := listener.Accept()
        if error != nil {
            fmt.Println(error)
        }
		fmt.Println("New client! Waiting for a name")
		client_name, _ := bufio.NewReader(connection).ReadString('\n')
		client_name = strings.TrimRight(client_name,"\r\n")
        client := &Client{socket: connection, data: make(chan string), name: client_name}
        manager.register <- client
        go manager.receive(client)
        go manager.send(client)
    }
}