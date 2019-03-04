package main

import (
	"os"
	"log"
	"text/template"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
// Run go run *.go
	foo()
}