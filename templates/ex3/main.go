package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "Pawel")

	if err != nil {
		log.Fatal(err)
	}

}
