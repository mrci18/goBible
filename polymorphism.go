package main

import (
	"fmt"
)

func (x person) who() {
	fmt.Printf("My name is %s %s and I am %d years old\n", x.first, x.last, x.age)
}

type human interface {
	who()
}

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}
func foo (h human) {
	h.who()
}

func main(){
	person1 := person {
		first: "Charles",
		last: "Ibo",
		age: 23,
	}

	person2 := person {
		first: "Stanley",
		last: "Ibo",
		age: 21,
	}

	sal := secretAgent {
		person: person{
			first: "Jenny",
			last: "Money",
			age: 19,
		},
		ltk: true,
	}

	person1.who()
	person2.who()
	foo(person1)
	foo(person2)
	foo(sal)

}