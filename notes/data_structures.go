package main

import (
	"fmt"
)

func main(){
	foo()
	bar()
	baz()
}

func foo(){
	list := []int{1,2,3}
	for i, v := range(list){
		fmt.Println(i, v)
	}
}

func bar(){
	xs := []string{"Hello", "World", "This", "Is", "Charles"}
	for i, v := range(xs){
		fmt.Println(i, v)
	}
}

func baz(){
	m := map[string]int{"Charles": 23, "Neil": 23, "Fonzy": 22}
	fmt.Println(m)

	for name, age := range(m){
		fmt.Println(name, age)
	}
}