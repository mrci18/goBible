package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	miToKm = 1.60934
)

func main(){
	number, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	//get the number of miles from the user
	// var number float64
	// fmt.Scanln(&number)

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")


	fmt.Printf("Miles: %.2f<br>\n", number)
	fmt.Printf("Kilometers: %.2f <br>\n", number*miToKm)

	fmt.Println("</body>")

	fmt.Println("</html>")
}