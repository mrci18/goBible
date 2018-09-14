package main

import (
	"fmt"
)

func main(){
	x := "ABCDE"



	fmt.Printf("%T\n", x)
	for _, v := range x {
		fmt.Println(v)
	}
}