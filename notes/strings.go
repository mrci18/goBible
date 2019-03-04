package main

import (
	"fmt"
	"strings"
	"strconv" //Converts string
)

func main(){
	// ToTitle prints all uppercase
	fmt.Println(strings.ToTitle("Hello World"))

	//Title prints First letters of word into uppercase
	fmt.Println(strings.Title("hello world"))

	//Replaces k with ky for the first two oink's
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))

	//Replace has no limit below 0. Replaces all oinks with moo
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -2))

	//Trim space
	fmt.Println(strings.TrimSpace("    test      "))

	//ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding 
	fmt.Println(strconv.ParseInt("AF0875", 16, 64))

	fmt.Println(strconv.FormatInt(123, 16))

	
}