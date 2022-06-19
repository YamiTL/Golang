package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Println("Enter your name: ")
	var name string 
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Yay, you can play!")
} else {
	fmt.Println("Boo, you cannot play!")
	return
}
}