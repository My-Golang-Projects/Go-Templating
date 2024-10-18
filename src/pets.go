package main

import (
	"os"
	"path/filepath"
	"text/template"
)

type Pet struct {
	Name   string
	Sex    string
	Age    string
	Intact bool
	Breed  string
}

func main() {
	dogs := []Pet{
		{Name: "Max", Sex: "Male", Age: "5", Intact: true, Breed: "Golden Retriever"},
		{Name: "Bella", Sex: "Female", Age: "3", Intact: false, Breed: "Labrador Retriever"},
		{Name: "Charlie", Sex: "Male", Age: "7", Intact: true, Breed: "German Shepherd"},
		{Name: "Lucy", Sex: "Female", Age: "2", Intact: false, Breed: "Beagle"},
		{Name: "Buddy", Sex: "Male", Age: "4", Intact: true, Breed: "Poodle"},
		{Name: "Molly", Sex: "Female", Age: "6", Intact: false, Breed: "Shih Tzu"},
		{Name: "Bailey", Sex: "Male", Age: "8", Intact: true, Breed: "Rottweiler"},
		{Name: "Sadie", Sex: "Female", Age: "1", Intact: false, Breed: "Chihuahua"},
	}

	var tmpfile = "src/pets.tmpl"

	// create a new template with the name of the file
	// then parse the file
	tmpl, err := template.New(filepath.Base(tmpfile)).ParseFiles(tmpfile)

	if err != nil {
		panic(err)
	}

	// Execute to print the finished report to the terminal, and also passing in our dogs slice
	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		panic(err)
	}
}
