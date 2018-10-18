package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Print(os.Args[0])
	fmt.Print(os.Args[1])
	str := fmt.Sprint(`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>
		`  + name + `
	</h1>
	</body>

	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	//Delete
	// fmt.Printf("str is of type ====== %T\n", str)
	// fmt.Printf("str is of type ====== %T\n", strings.NewReader(str))
}

