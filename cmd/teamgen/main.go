//  teamgen/main.go

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

const supp4 = `{{ .Name }} [{{ .Gender }}] {{ .UPPs }} Age: {{ .Age }} {{ .Species }}
{{ .Terms }} terms {{ title .Career }}
{{ .SkillString }}
`

const brp = `{{ .Name }} [{{ .Gender }}] Age: {{ .Age }} {{ .Species }} {{ .Terms }} terms {{ title .Career }}
{{ .UPPs }}
{{ .SkillString }}
`

var funcMap = template.FuncMap{
	"title": strings.Title,
}

var outstring bytes.Buffer

func outstringer(p person.Person, tmpl string) string {
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
}

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

	//'game,job,career,terms,gender,first name, middle name, last name, suffix'")
	career := flag.String("career", "", "Career or Branch")
	game := flag.String("game", "2d6", "Game version")
	gender := flag.String("gender", "", "F or M, default random")
	job := flag.String("job", "", "Job")
	lastName := flag.String("lastName", "", "Last name")
	terms := flag.String("terms", "", "Number of terms, random 1-5")

	listOptions := flag.Bool("list", false, "List Career and Job options")
	flag.Parse()

	options["datadir"] = datadir
	careerFile := path.Join(datadir, "careers.txt")
	jobFile := path.Join(datadir, "jobs.txt")
	if *listOptions {
		fmt.Println(datamine.ListOptions(careerFile, jobFile))
		os.Exit(0)
	}

	options["gender"] = strings.ToUpper(*gender)
	options["terms"] = *terms
	options["career"] = *career
	options["job"] = *job
	options["lastName"] = *lastName
	options["db_name"] = path.Join(datadir, "names.db")
	options["game"] = *game
	p := person.MakePerson(options)

	result := ""
	if *game == "brp" {
		result = outstringer(p, brp)
	} else {
		result = outstringer(p, supp4)
	}
	fmt.Println(result)
}
