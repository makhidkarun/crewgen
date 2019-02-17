
package main

import (
  //"fmt"
  "html/template"
  "log"
  "net/http"
  "strings"

  "github.com/makhidkarun/crewgen/pkg/person"
)

var templateDir string = "web"
type Crew []person.Person
type Ship struct {
  ShipName  string
  Crew      []person.Person
}

func buildCrew() Crew {
  var crew Crew
  var options = make(map[string]string)
  options["terms"] = "0"
  options["gender"] = ""
  options["db_name"] = "data/names.db"
  captain := person.MakePerson(options)
  pilot := person.MakePerson(options)
  crew = append(crew, captain)
  crew = append(crew, pilot)
  return crew
}

func showCrew (w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  sN                := r.Form["shipName"]
  shipName          := strings.Join(sN, "")
  crew              := buildCrew() 
  ship              := Ship{ShipName: shipName, Crew: crew}
  htmlOpenTemplate  := templateDir + "/" + "formTemplate.html"
  htmlCloseTemplate := templateDir + "/" + "formTemplate.html"
  formTemplate      := templateDir + "/" + "formTemplate.html"
  crewTemplate      := templateDir + "/" + "crewTemplate.html"
  personTemplate    := templateDir + "/" + "personTemplate.html"
  t,err := template.ParseFiles(htmlOpenTemplate,
      htmlCloseTemplate,
      formTemplate,
      crewTemplate,
      personTemplate) 
  if err != nil {
    log.Fatal("Got an error %s.", err)
  }
  t.Execute(w, ship)
}

func recruitCrew (w http.ResponseWriter, r *http.Request) {
  formTitle           := "Recruiting Your Crew"
  htmlOpenTemplate    := templateDir + "/" + "formTemplate.html"
  htmlCloseTemplate   := templateDir + "/" + "formTemplate.html"
  formTemplate        := templateDir + "/" + "formTemplate.html"
  t,_                 := template.ParseFiles( htmlOpenTemplate, formTemplate, htmlCloseTemplate)
  t.Execute(w, formTitle)
}

func main(){
  server := http.Server {
    Addr: "127.0.0.1:8080",
  }

  http.HandleFunc("/recruit", recruitCrew)
  http.HandleFunc("/show", showCrew)
  log.Println("Starting server.")
  log.Fatal(server.ListenAndServe())

}

