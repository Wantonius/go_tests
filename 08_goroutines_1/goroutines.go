package main

import (  
    "fmt"
    "time"
)


func hello() {  
    fmt.Println("Hello world goroutine")
}

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
	fmt.Printf("numbers thread terminates");
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
	fmt.Printf("alphabets thread terminates");
}


func main() { 

	//go command will create another execution context or a thread. P
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("In main function")
	
	//slightly more involved example
	
	go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("... main ends")
}