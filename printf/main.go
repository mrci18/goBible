package main

import "fmt"

func main() {
	var brand = "Google"
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	//Quote the string,
	//You can use as many verbs
	ops, ok, fail := 2350, 543, 443

	fmt.Printf("%q\n", brand)

	fmt.Printf(
		"total: %d success: %d / %d\n",
		ops, ok, fail,
	)
	//Prints type
	fmt.Printf("%T\n", ops)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions of km\n", distance)
	fmt.Printf("Orbital period: %f\n", orbital)
	fmt.Printf("Does %v have life? %t\n", planet, hasLife)

	//Positional args
	fmt.Printf("%v is %v away. Think %[2]v kms\n Planet %[1]v", planet, distance)

	//Precision
	fmt.Printf("Orbital period: %.0f days", orbital)
}
