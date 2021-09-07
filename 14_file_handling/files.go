package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFile(filename, text string) {

	fmt.Printf("Writing to a file in Go lang\n")
	

	file, err := os.Create(filename)
	
	if err != nil {
		fmt.Println("Can't create file %s. Reason %s\n",filename,err)
		os.Exit(1)
	}
	
	defer file.Close()
	

	len, err := file.WriteString(text)
	if err != nil {
		fmt.Println("Can't write to file %s, Reason %s\n",filename,err)
		os.Exit(1)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")
	
	data, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Println("Can't read from file %s. Reason %s\n",filename,err)
		os.Exit(1)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	// user input for filename
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	// user input for file content
	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	
	// file is created and read
	CreateFile(filename, input)
	ReadFile(filename)
}
