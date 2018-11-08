// package main

// import (
// 	"log"
// 	"os"
// 	"text/template"
// )

// func main() {
// 	tpl, err := template.ParseFiles("tpl.gohtml")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = tpl.Execute(os.Stdout, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

package main 

import (
	"log"
	"text/template"
	"os"
)
func main (){
	tpl, err := template.ParseFiles("one.txt", "two.txt")
	if err != nil {
		log.Fatal("whooops", err)
	}

	tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	tpl.ExecuteTemplate(os.Stdout, "two.txt", "Charles")
}