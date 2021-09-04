package main

import (
    "errors"
    "fmt"
)

//errors are very nice for additional return values so that you can return a proper value AND an error at the same time.
//This behavior is often needed in such functions that can create errors but also partially succeed. Such as reading from or
//writing to a source. You partially manage the feat but then an error occurs. In older languages such as C when this
//sort of behavior is required you first return the succeeding part and then because it was not fully successful on
//next call you return the error. In go we can return both the partial success and the resulting error at the same time.

func notDog(animal string) (string,error) {
	if(animal == "dog") {
		return "", errors.New("Not working with a dog!")
	}
	return animal+" is not a dog, will work fine",nil
}

func main() {

	animals := []string{"cat","dog"}
	
	for _, a := range animals {
		result, err := notDog(a);
		if(err != nil) {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}