package main

import (
	"log"
	"os"
	"text/template"
)

// 1. Parse your templates, preferrably all at once in glob
// 2. Then execute templates

func main() {

	// from package template we are going to parse files, and we provide file name.
	// ParseFiles can take multiple strings, and returned pointers to parsed templates and errors.
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating index.html", err)
	}
	defer nf.Close()

	// we are asking for result of file parsing, using pointer 'nil'
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//
	// Now we will parse multiple files at once.
	tpl, err = tpl.ParseFiles("one.gohtml", "two.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// now we will execute templates with template names by using ExecuteTemplate
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//
	// We can instead parse a whole folder of *.gohtml using ParseGlob
	// Protip, parse all of your templates in one go for performance
	tplGlob, errGlob := template.ParseGlob("templates/*.gohtml")
	if errGlob != nil {
		log.Fatalln(errGlob)
	}

	// now we execute to see the first template in the glob
	errGlob = tplGlob.Execute(os.Stdout, nil)
	if errGlob != nil {
		log.Fatalln(errGlob)
	}
	// execute a specific template by name in the glob
	errGlob = tplGlob.ExecuteTemplate(os.Stdout, "charlie.gohtml", nil)
	if errGlob != nil {
		log.Fatalln(errGlob)
	}
}
