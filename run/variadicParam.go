package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2,3,4,}
	foo(xi...)
}

func foo(i ...int){
	fmt.Println(i)
}