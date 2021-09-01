package main

import "fmt"

func main() {
   var b int = 15
   var a int
   var t int = 10;
   
   //note that the initialization does not cover the whole array length
   
   numbers := [6]int{1, 2, 3, 5} 

   /* for loop execution */
   for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
   }
   for a < b {
      a++
      fmt.Printf("value of a: %d\n", a)
   }
  
	//i,x for index and value. Replace index with blank if needed "for _,x" 
   
   for i,x:= range numbers {
      fmt.Printf("value of x = %d at %d\n", x,i)
   }

	//If has no parentheses around conditions in Go, but that the braces are required.
	
	if t > 5 {
		fmt.Printf("%d is bigger than 5\n",t)
	}
	
	if t < 5 {
		fmt.Printf("%d is smaller than 5\n",t)
	} else {
		fmt.Printf("%d is still bigger than 5\n",t)
	}
	
	//Switch is more open than in other languages. You can compare things like types not just numbers or strings.
	
	i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
	
	whatAmI := func(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
	
}
