package main

import (
	"os"
	"path/filepath"
	"strings"
	"html/template"
)

func main_html() {
	dogs := samplePets()

	funcMap := template.FuncMap{
		"dec":     func(i, j int) int { return i - j },
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}

	var tmpfile = "src/petsHtml.tmpl"

	tmpl, err := template.New(filepath.Base(tmpfile)).Funcs(funcMap).ParseFiles(tmpfile)

	if err != nil {
		panic(err)
	}

	var f *os.File

	f, err = os.Create("pets.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, dogs)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
