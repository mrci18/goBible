package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "about.gohtml", nil)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	f, err := os.Create("main.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f2, err := os.Create("about.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()


	err = tpl.ExecuteTemplate(f, "main.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(f, "about.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}