package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpl := template.New("test")

	err := tmpl.Execute(os.Stdout, nil)
	check(err)
}
