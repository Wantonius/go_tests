package main

import "fmt"

func main() {

	//This is an array in go. It has a _specific_ length and will be initialized to zero-values (ie. zero-valued) unless explicitly initialized to something else.
	var myArray [6]int
	
	fmt.Println("myArray:",myArray)
	fmt.Println("myArray length:",len(myArray))
	
	myArray[3] = 50
	fmt.Println("myArray after set:",myArray)
	
	myInitializedArray := [3]int{0,1,2}
		
	fmt.Println("Initialized Array:",myInitializedArray)
	
	fmt.Println("---Slices---")

	//This is a slice. Far more useful interface for sequences in go. Slices do not have a set length only set type they contain. Slice support additional methods
	//when compared to array. You will mostly use slices.
	
	var mySlice []int //this does not allocate memory for the slice
	
	myAllocatedSlice := make([]int,10) //this allocates memory for a slice of 10 with zeroes
	
	//Effectively slice is a dynamically allocated array. Far more useful than static one.
	
	fmt.Println("mySlice:",mySlice)
	fmt.Println("mySlice length:",len(mySlice))
	
	fmt.Println("myAllocatedSlice:",myAllocatedSlice)
	fmt.Println("myAllocatedSlice length:",len(myAllocatedSlice))
	
	// Since we did not reserve memory for mySlice, following line WILL crash the program 
	// mySlice[0] = 0
	// You should use append in those cases. Or use make and reserve the memory in creation.
	mySlice = append(mySlice,0)
	
	// When appending slices to slices you need to use the three dots operator '...'. It will tell go that the slice will contain same type and can be appended. Otherwise
	// you will get an error saying that the types do not match.
	
	mySlice = append(mySlice,[]int{10,100}...)
	
	fmt.Println("mySlice again:",mySlice)
	fmt.Println("mySlice length:",len(mySlice))
	
	copiedSlice := make([]int,len(mySlice))
	
	copy(copiedSlice,mySlice)
	
	fmt.Println("CopiedSlice:",copiedSlice)
	
	partialSlice := mySlice[1:3]
	
	fmt.Println("PartialSlice:",partialSlice)
	
	fmt.Println("---Maps---")
	
	//Maps are go’s built-in associative data type (sometimes called hashes or dicts in other languages).
	//To create an empty map, use the builtin make: make(map[key-type]val-type)
	
	intStringMap := make(map[int]string)
	stringIntMap := make(map[string]int)
	
	intStringMap[1] = "One"
	intStringMap[2] = "Two"
	
	stringIntMap["one"] = 1
	stringIntMap["two"] = 1
	
	fmt.Println("intStringMap",intStringMap)
	fmt.Println("stringIntMap",stringIntMap)
	fmt.Println("Value at 1:",intStringMap[1])
	fmt.Println("Value at one:",stringIntMap["one"])
	
	//use delete to remove by key
	
	delete(stringIntMap,"two");
	fmt.Println("stringIntMap again:",stringIntMap)
	
	//initalize map for init values
	
	initializedMap := map[int]string{1:"one",2:"two"}
	fmt.Println("initializedMap",initializedMap);
}