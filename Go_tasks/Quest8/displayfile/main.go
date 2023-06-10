package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	filename := args[0]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf(string(content))
}
