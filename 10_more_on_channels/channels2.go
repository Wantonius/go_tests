package main

import (  
    "fmt"
    "time"
)

func worker(ch chan string, s time.Duration) {  
    time.Sleep(s * time.Millisecond)
    ch <- "worker done"
}

func main() {  
    channel := make(chan string)
	channel2 := make(chan string)
    go worker(channel,6500)
	go worker(channel2,3500)
	
	//Select waits for multiple channel sources and works like switch-case. Default case is for situations where none
	//of the sources have input to read. It is also used for so called non-blocking operations. 
	//Note the way to break out of the named loop. Break will not work properly without it in this case.
	L:
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-channel:
            fmt.Println("Worker 1 says: ", v)
            break L
		case v := <-channel2:
			fmt.Println("Worker 2 says: ", v)
        default:
            fmt.Println("no value received")
        }
    }
	fmt.Println("--- Moving on to closing channels ---");

    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
			fmt.Println("Worker: Waiting for more jobs!")
			//The additional return value "more" will be false when the channels gets closed. So this is a way to monitor the channels status.
			//We will close the jobs channel 
            j, more := <-jobs
            if more {
                fmt.Println("Worker: Received job", j)
            } else {
                fmt.Println("Worker: Received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        fmt.Println("Main: Sending another job")
		jobs <- j
        fmt.Println("Main:Sent job", j)
		time.Sleep(1000 * time.Millisecond)
    }
    close(jobs)
    fmt.Println("Main:Sent all jobs")

    <-done
	
	fmt.Println("-- Ranging over buffered channels ---")

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

	fmt.Println("Main: closed the queue channel")
	//Messages already in the channels WILL BE delivered. 

    for elem := range queue {
        fmt.Printf("Received element %s from the closed channel\n",elem)
    }
}