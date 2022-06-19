package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	name := "Yami"
	age := 29
	// another possibility: var name string = "Yami" 
	fmt.Printf("Hello, %v, how are you doing?", name)
	
	fmt.Printf(" %v, please confirm you are %v years old", name, age)
}