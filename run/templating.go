package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(&x)


	//Pointer to an int, *type
	y := &x
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//No longer want the refenrence, want the value 
	fmt.Println(*y)

	*y = 43

	fmt.Println(x)
	
}

