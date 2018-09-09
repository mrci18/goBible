package main

//Blank identifier _
import (
	"fmt"
)

//Assign variable with var
var x = 5

//Assign variable with const
const (
	r = 5
	t = 7
)

func main(){
	//Assign variable shorthand can't change type after this way
	y := 10

	y =11
	fmt.Println(y)
	fmt.Println(x)
	charlesTest()

}

func charlesTest(){
	fmt.Println(x)
}