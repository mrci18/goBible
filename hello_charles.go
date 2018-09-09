package main

import "fmt"

func main() {
	// x := "Charles"
	// fmt.Println("Hello ",x)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello my name is", name)
}