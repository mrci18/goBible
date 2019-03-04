package main

import "fmt"

func main(){
	// i:=0
	// for i <10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	//Refers to index and prints index
	for x:=range "test"{
		fmt.Println(x)
	}

	//Prints string val
	str := "test"
	for i:= range str {
		fmt.Println(str[i])
	}

	//You can use blank identifier
	for _, c := range "test" {
		fmt.Println(c)
	}

	condt := true
	for condt {
		fmt.Println("Hello")
		condt = false
	}

	//If statement check if even or odd
	//else if {}
	length := 10
	for i:=0; i <= length; i++{
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	for i := 0; i <= 10; i++ {
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		default: fmt.Println("Above 2 or below 0")
		}
	}

}

