package main 

import "fmt"

func main(){
	var x string = "hello_world"
	const name = "Charles"
	var (
		t,s= "Tony", "Snell"
	)

	fmt.Println(&x)
	fmt.Println(t,s)
}