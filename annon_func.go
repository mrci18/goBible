package main

import (
	"fmt"
)

func main() {
	defer foo()
	//annonymous func
	func(x int){
		fmt.Println(x)
	}(4)
}

func foo(){
	fmt.Println("foo")

}