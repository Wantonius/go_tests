package main

import "fmt"

//basic functions. Note that the variable name is before type and the return value is last.  

func add(a int, b int) int {
	return a+b;
}

func substract(a int, b int) int {
	return a-b;
}

//multiple return values 

func math_action(a int, b int, action string)(int,string) {
	if action == "add" {
		return a+b,"added"
	} 
	if action == "substract" {
		return a-b,"substracted"
	}
	return 0,"unknown action"

}

//variadic functions

func sum(vals ...int) int {
	total := 0
	for _,num := range vals {
		total += num
	}
	return total
}

//closures or anonymous functions returned by function. In a sense these are factories.

func sequence() func() int {
	i := 0;
	return func() int {
		i++;
		return i;
	}
}

func main() {

	c := add(10,20)
	d := substract(20,10)
	fmt.Println("10+20 = ",c)
	fmt.Println("20-10 = ",d)
	
	e,act := math_action(20,10,"add")
	f,act2 := math_action(20,10,"substract")
	
	fmt.Printf("20 %s to 10 leads to %d\n",act,e)
	fmt.Printf("10 %s from 20 leads to %d\n",act2,f)

	t := []int{1,2,3,4,5,6,7,8,9}
	
	fmt.Println("Adding together these numbers", t)
	fmt.Printf("leads to %d\n",sum(t...))
	
	nextInt := sequence()
	
	fmt.Println("Ints in sequence")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := sequence()
	
	fmt.Println("Let's restart")
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}