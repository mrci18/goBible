package main

import (
	"fmt"
)

	type person struct {
		first string
		age int
		saying string
	}

	type secretagent struct {
		person
		ltk bool
	}

	//func receiver identifier(params) returns (code)
	func (p person) speak(){
		fmt.Println(p.first, "says", p.saying)
	}

	func (sa secretagent) speak(){
		fmt.Println(sa.first, "says", sa.saying)
	}

	type human interface {
		speak()
	}

	func foo (h human){
		h.speak()
	}


var c int
var d = 42
func main(){
	a:= "James"
	fmt.Println(a)
	// Prints type of a
	fmt.Printf("%T\n", a)

	//slice Aggregate composite data structure
	xi := []int{2,3,4,5,}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}

	//map AGGREGATE or COMPOSITE data structure
	f := map[string]int{"James": 32, "Jenny": 24, "Greg": 25,}
	fmt.Println(f)

	for k, v := range f {
		fmt.Println(k,v)
	}

	//struct, composite literal type goes first(person)
	p1 := person {
		first: "Charles",
		age: 23,
		saying: "I like to wrestle",
	}

	p2 := person {
		first: "James",
		age: 27,
		saying: "I like to play basketball",

	}

	people := []person{p1,p2}

	for k2, v := range people {
		fmt.Println(k2,v)
	}

	for j:=0; j<10; j++ {
		fmt.Println(j)
	}

	sa1 := secretagent {
		person: person {
			first: "Jeff",
			age: 21,
			saying: "I'm a secret agent",
		},
		ltk: true,

	}
	fmt.Println(sa1)
	fmt.Println(sa1.first)
	fmt.Println(sa1.person.first)
	foo(p1)
}