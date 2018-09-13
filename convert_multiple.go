package main

import (
	"fmt"
)

const (
	miToKm = 1.60934
	pToKg = 0.453592
	divider = "+--------------------------------------------------------+"
)

func main(){
	fmt.Println("Choose an option ")
	fmt.Println("1: Miles to kilometers: ")
	fmt.Println("2: Fahrenheit to Celcius")
	fmt.Println("3: Pounds to kilograms")

	var (
		option int
		number float64
	)

	fmt.Scanln(&option)

	// fmt.Println("Enter a number to convert: ")
	// fmt.Scanf("%f", &number)

	fmt.Println(divider)
	switch option {
	case 1: 
		fmt.Scanln(&number)
		fmt.Printf("| Miles: %15.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("| Kilometers: %10.2f |\n", number*miToKm)
	case 2: 
		fmt.Scanln(&number)
		fmt.Printf("| Fahrenheit: %.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("Celsius: %13.2f |\n", (number-32)*5/9)
	case 3: 
		fmt.Scanln(&number)
		fmt.Printf("| Pounds: %.2f | \n", number)
		fmt.Println(divider)
		fmt.Printf("| Kilos: %.2f | \n", number*pToKg)
	}
	fmt.Println(divider)
}

