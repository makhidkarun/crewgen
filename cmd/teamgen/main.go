//  cli/main.go

package main

import (
	"bytes"
  "flag"
	"fmt"
	//"os"
	"text/template"

	"github.com/makhidkarun/crewgen/pkg/person"
)

const supp4 = `{{ .Name }} [{{ .Gender }}] {{ .UPPs }} Age: {{ .Age }} {{ .Species }}
{{ .Terms }} terms {{ .Career }}
{{ .SkillString }}
`

func main() {
	var options = make(map[string]string)
	var outstring bytes.Buffer

  gender := flag.String("gender", "", "F or M, default random")
  flag.Parse()
  //fmt.Printf("gender is %s.\n", *gender)
	//options["terms"] = "2"
	//options["gender"] = "F"
	options["job"] = "pilot"
	options["db_name"] = "data/names.db"

  options["gender"] = *gender
	p := person.MakePerson(options)
	tmpl, err := template.New("supp4").Parse(supp4)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&outstring, p)
	if err != nil {
		panic(err)
	}
	result := outstring.String()
	fmt.Println(result)
}
