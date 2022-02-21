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
	terms := flag.String("terms", "0", "Number of terms, default 1-4")
	career := flag.String("career", "", "Career or Branch")
	flag.Parse()
	options["gender"] = *gender
	options["terms"] = *terms
	options["career"] = *career
	options["job"] = "pilot"
	options["db_name"] = "data/names.db"

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
