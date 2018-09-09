package main

import (
	"fmt"
)
//Unsigned integer. Same thing as int but can't be negative
var x uint = 5
var c,i = 1,2
var n = "Hello \n"
var hello = "Hello" + "World" + fmt.Sprint(x)
var str = `Multiple 

lines`
func main(){
	fmt.Println(x)
	fmt.Println(hello)
	fmt.Println(len(str))

	fmt.Println(hello[2])
}

