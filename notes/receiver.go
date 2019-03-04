package main

import (
	"fmt"
)

func (x car) ads() {
	fmt.Printf("The %d %s %s is the best selling car of %d\n", x.year, x.make, x.model, x.year) 
}

type car struct {
	make string
	model string
	year int
}

func main(){
	car1 := car {
		make: "honda",
		model: "civic",
		year: 2018,
	}

	car2 := car {
		make: "toyota",
		model: "corolla",
		year: 2017,
	}
	car1.ads()
	car2.ads()

}