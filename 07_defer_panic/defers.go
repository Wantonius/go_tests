package main

import "fmt"

//defer is another nice little feature of golang. It allows you to call a function beforehand and defer the execution to the end to the current function.
//This behavior is extremely nice for closing files and sockets and such. Also remembering to unlock mutexes, release resources etc.

func willExecuteLast(greet string) {
	fmt.Printf("Goodbye %s, I was deferred to be last in the calling function\n",greet)
}

func callsAdditionalDefer(greet string) {
	defer willExecuteLast(greet)
	fmt.Println("I will be before the first goodbye")
}

func helloGreeting(greet string) {
	fmt.Printf("Hello %s, I will execute first\n",greet)
}

//A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.

func panics() {
	panic("calamity ensues!")
}

func main() {

	//recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
	//r will be the panic message.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("It panicked but we recovered. Error:%s\n",r)
		}
	}()

	defer panics();

	defer fmt.Println("Next we panic and recover!");
	
	defer willExecuteLast("John")
	defer callsAdditionalDefer("Johnny")
	
	fmt.Println("First we test defer")
	helloGreeting("John")


	

	
}