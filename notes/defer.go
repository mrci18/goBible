package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo(){
	fmt.Println("foo")
	defer one()
	two()
}

func bar(){
	fmt.Println("bar")
}

func one () {
	fmt.Println("One")
}

func two () {
	fmt.Println("two")
}