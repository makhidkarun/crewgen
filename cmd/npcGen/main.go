//  npcGen/main.go

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/makhidkarun/crewgen/pkg/datamine"
	"github.com/makhidkarun/crewgen/pkg/person"
)

// var funcMap changes the title
var funcMap = template.FuncMap{
	"title": strings.Title,
}

// var outstring is the bytes buffer for the printable output.
var outstring bytes.Buffer

// func outstringer takes a person and a template, and returns a string.
func outstringer(p *person.Person, options map[string]string) string {
	/*
		t := template.New("t")
		t, err := t.Funcs(funcMap).Parse(tmpl)
		if err != nil {
			panic(err)
		}
		err = t.Execute(&outstring, p)
		if err != nil {
			panic(err)
		}
		result := outstring.String()
		return result
	*/
	var b bytes.Buffer
	//templateDir := "templates"
	personTFile := path.Join(options["templatedir"], "supp4.tmpl")
	_, err := os.Stat(personTFile)
	if err != nil {
		fmt.Printf("Error:  %q", err)
	}
	t, err := template.ParseFiles(personTFile)
	if err != nil {
		fmt.Println(err)
	}
	// Fix upper case before output
	// better to go back to the funcmap?
	p.Career = strings.Title(p.Career)
	err = t.ExecuteTemplate(&b, "person", p)
	result := b.String()
	return result
}

// whine is a rudimentary error handler
func whine(err error) {
	if err != nil {
		fmt.Printf("Error: %q", err)
	}
}

func main() {

	var options = make(map[string]string)

	exe, err := os.Executable()
	whine(err)
	exedir := path.Dir(exe)
	datadir := path.Join(exedir, "data")
	templatedir := path.Join(exedir, "templates")

	career := flag.String("career", "", "Career or Branch")
	game := flag.String("game", "2d6", "Game version")
	gender := flag.String("gender", "", "F or M, default random")
	job := flag.String("job", "", "Job")
	lastName := flag.String("lastName", "", "Last name")
	terms := flag.String("terms", "", "Number of terms, random 1-5")

	listOptions := flag.Bool("list", false, "List Career and Job options")
	flag.Parse()

	options["datadir"] = datadir
	options["careerFile"] = path.Join(datadir, "careers.txt")
	options["jobFile"] = path.Join(datadir, "jobs.txt")
	if *listOptions {
		fmt.Println(datamine.ListOptions(options["careerFile"], options["jobFile"]))
		os.Exit(0)
	}

	options["gender"] = strings.ToUpper(*gender)
	options["terms"] = *terms
	options["career"] = *career
	options["job"] = *job
	options["lastName"] = *lastName
	options["game"] = *game
	options["templatedir"] = templatedir
	p := person.MakePerson(options)

	result := outstringer(&p, options)
	fmt.Println(result)
}
