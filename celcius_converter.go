package main

import "fmt"


func main() {
	var celcius float64

	fmt.Print("Enter Celcius: ")
	_, err:= fmt.Scanf("%v", &celcius)
	fmt.Println(err)
	fahrenheit := celcius * 1.8 + 32

	fmt.Println("Fahrenheit", fahrenheit)

}