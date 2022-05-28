package main

import (
	"fmt"
	"io/ioutil"
)

func main (){
	// write data into a file
	fmt.Println("Writing data into a file")
	writeFile("Welcome to go")
	readFile()

	// read data from a file
	fmt.Println("Reading data from a file")
	readFile()
}

func writeFile(message string){
	bytes := []byte(message)
	ioutil.WriteFile("test.txt", bytes, 0644)
	fmt.Println("File written successfully")
}

func readFile(){
	data, _ :=	ioutil.ReadFile("test.txt")
	fmt.Println("File content: ", string(data))
}