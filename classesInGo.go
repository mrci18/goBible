package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
	sayings []string
}

func main(){
	p1 := person{
		first: "Charles",
		last: "Ibo",
		age: 23,
		sayings: []string{"Hell world", "I like golang"},
	}
	fmt.Println(p1)

	p2 := person{
		first: "James",
		last: "Bonds",
		age: 32,
		sayings: []string{"Hello world", "I like python"},
	}	
	fmt.Println(p2)

	xp := []person{p1,p2}

	for i, v := range xp {
		fmt.Println(i,v)
		fmt.Println(v.first)
		for j,w := range v.sayings {
			fmt.Println(j, w)
		}
	}

}
