package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"strings"
)

func main() {
	//Create string and print
	name := "Charles"
	str := `html here` + name + `more html`
	fmt.Println(str)

	//Create string and print
	s := fmt.Sprint(`more` + name + `1 more`)
	fmt.Println(s)

	//Create file 
	//io.Copy() to the file
	nf, err := os.Create("new_file.txt")
	if err != nil{
		log.Fatal("Error - ", err)
	}
	io.Copy(nf, strings.NewReader(s))

	//Create file
	nf2, err := os.Create("new_file2.txt")
	if err != nil{
		log.Fatal("Error - ", err)
	}

	//Writes string to a file
	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("Error - ", err)
	}
	
	fmt.Println("bytes written", n)
}
