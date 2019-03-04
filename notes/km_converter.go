package main

import "fmt"

const milesToKM = 1.60mv 934

func main() {


	var miles float64
	fmt.Print("Enter miles: ")
	fmt.Scanf("%f", &miles)

	kilometers := miles * milesToKM

	fmt.Println("Kilometers",kilometers)
}