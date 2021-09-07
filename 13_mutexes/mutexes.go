package main  
import (  
    "fmt"
    "sync"
    )
	
var x  = 0  //a global variable represents the critical section

func increment(wg *sync.WaitGroup) {  
    x = x + 1 //every goroutine tries to add to global variable causing a race condition
    wg.Done()
}

func fixed_increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock() //We lock the access to the critical section before attempting to change the global state
    x = x + 1 
    m.Unlock() //Remember to unlock ALWAYS. Use defer liberally with these.
    wg.Done()   
}

func main() {  
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {	// lets create 1000 goroutines 
        w.Add(1)        
        go increment(&w)
    }
    w.Wait()
	fmt.Println("Variable SHOULD come up to a 1000. It almost never will because of race conditions")
    fmt.Println("final value of x", x)
	
	fmt.Println("So lets fix this")
	
	x = 0
	
	var m sync.Mutex
	
	for i := 0; i < 1000; i++ {	// lets create 1000 more goroutines 
        w.Add(1)        
        go fixed_increment(&w,&m)
    }
	
	w.Wait()
	
	fmt.Println("final value of x", x)
}