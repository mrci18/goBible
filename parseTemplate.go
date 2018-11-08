package main

import (
	"text/template"
	"log"
	"os"
)

func main(){
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index2.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}