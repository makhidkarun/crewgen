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

var funcMap = template.FuncMap{
	"title": strings.Title,
}

func whine(err error) {
	if err != nil {
		fmt.Printf("Error: %q", err)
	}
}

func main() {

	var options = make(map[string]string)
	var outstring bytes.Buffer

	exe, err := os.Executable()
	whine(err)
	exedir := path.Dir(exe)
	datadir := path.Join(exedir, "data")

	gender := flag.String("gender", "", "F or M, default random")
	terms := flag.String("terms", "0", "Number of terms, random 1-5")
	career := flag.String("career", "", "Career or Branch")
	job := flag.String("job", "", "Job")
	listOptions := flag.Bool("list", false, "List Career and Job options")
	flag.Parse()

	options["datadir"] = datadir
	careerFile := path.Join(datadir, "careers.txt")
	jobFile := path.Join(datadir, "jobs.txt")
	if *listOptions {
		fmt.Println(datamine.ListOptions(careerFile, jobFile))
		os.Exit(0)
	}

	options["gender"] = *gender
	options["terms"] = *terms
	options["career"] = *career
	options["job"] = *job
	options["db_name"] = path.Join(datadir, "names.db")
	p := person.MakePerson(options)
	tmpl := template.New("supp4")
	tmpl, tErr := tmpl.Funcs(funcMap).Parse(supp4)
	if tErr != nil {
		panic(err)
	}
	err = tmpl.Execute(&outstring, p)
	if err != nil {
		panic(err)
	}
	result := outstring.String()
	fmt.Println(result)
}
