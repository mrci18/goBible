package main

import (
	"log"
	"os"
	"html/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	jesus := sage {
		Name: "Jesis",
		Motto: "Love all",
	}

	sages := []sage{buddha, jesus}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}