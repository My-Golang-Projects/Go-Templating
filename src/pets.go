package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Pet struct {
	Name   string
	Sex    string
	Age    string
	Intact bool
	Breed  []string
}

func samplePets() []Pet {
	return []Pet{
		{Name: "Max", Sex: "Male", Age: "5", Intact: true, Breed: []string{"Golden Retriever"}},
		{Name: "Bella", Sex: "Female", Age: "3", Intact: false, Breed: []string{"Labrador Retriever"}},
		{Name: "Charlie", Sex: "Male", Age: "7", Intact: true, Breed: []string{"German Shepherd"}},
		{Name: "Lucy", Sex: "Female", Age: "2", Intact: false, Breed: []string{"Beagle"}},
		{Name: "Buddy", Sex: "Male", Age: "4", Intact: true, Breed: []string{"Poodle"}},
		{Name: "Molly", Sex: "Female", Age: "6", Intact: false, Breed: []string{"Shih Tzu"}},
		{Name: "Bailey", Sex: "Male", Age: "8", Intact: true, Breed: []string{"Rottweiler"}},
		{Name: "Sadie", Sex: "Female", Age: "1", Intact: false, Breed: []string{"Chihuahua"}},
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  []string{"German Shepherd", "Pitbull"},
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  []string{"German Shepherd", "Border Collie"},
		},
		{
			Name:   "<script>alert(\"Gotcha!\");</script>Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  []string{"German Shepherd", "Border Collie"},
		},
	}
}

func main() {
	dogs := samplePets()

	main_html()
	funcMap := template.FuncMap{
		"dec":     func(i, j int) int { return i - j },
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}

	var tmpfile = "src/pets.tmpl"

	// create a new template with the name of the file
	// then parse the file
	tmpl, err := template.New(filepath.Base(tmpfile)).Funcs(funcMap).ParseFiles(tmpfile)

	if err != nil {
		panic(err)
	}

	// Execute to print the finished report to the terminal, and also passing in our dogs slice
	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		panic(err)
	}
}
