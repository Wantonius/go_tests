package main


import (
    "fmt"
    "time"
)


//function for synchronization

func worker(done chan bool) {
    fmt.Println("Worker: Lets do some work!")
    time.Sleep(3*time.Second)
    fmt.Println("Worker: And we are done")

	fmt.Println("Worker: Sending true to channel indicating that we are done")
	done <- true
}

func main() {

    messages := make(chan string)

	//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	//By default sends and receives block until both the sender and receiver are ready.
	
	fmt.Println("---Basic channel---")

    go func() {
		fmt.Println("Pinger:Pinging the main")
		messages <- "ping" 
	}()

	fmt.Println("Main:Reading the channel")
    msg := <-messages
    fmt.Println(msg)
	time.Sleep(2*time.Second)
	//By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. 
	//You can create buffered channels like so
	
	fmt.Println("---Buffered channel---")
	
	buffered := make(chan string, 2)
	
	buffered <- "buffered"
    buffered <- "channel"
	
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	time.Sleep(2*time.Second)
	//We can use channels to synchronize execution across goroutines.
	fmt.Println("---Channel synchro---")
	
	done := make(chan bool, 1)
    go worker(done)
	fmt.Println("Main: Wait for the worker to complete")
	//we can optionally get the channel message but since we are only using this as a notification that our worker is done it is not necessary
	<-done
	fmt.Println("Main: Worker is done. Main exiting")
}