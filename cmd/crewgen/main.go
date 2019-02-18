
package main

import (
  "html/template"
  "log"
  "net/http"
  "path/filepath"
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
  sN         := r.Form["shipName"]
  shipName   := strings.Join(sN, "")
  crew       := buildCrew() 
  ship       := Ship{ShipName: shipName, Crew: crew}
  layoutST   := filepath.Join(templateDir, "layoutS.tmpl")
  htmlOpenT  := filepath.Join(templateDir, "htmlOpen.tmpl")
  htmlCloseT := filepath.Join(templateDir, "htmlClose.tmpl")
  formT      := filepath.Join(templateDir, "form.tmpl")
  crewT      := filepath.Join(templateDir, "crew.tmpl")
  personT    := filepath.Join(templateDir, "person.tmpl")
  t,err      := template.ParseFiles(layoutST, htmlOpenT, htmlCloseT, formT, crewT, personT) 
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
    return
  }
  if err := t.ExecuteTemplate(w, "layoutS", ship); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

func recruitCrew (w http.ResponseWriter, r *http.Request) {
  formTitle    := "Recruiting Your Crew"
  layoutT      := filepath.Join(templateDir, "layout.tmpl")
  htmlOpenT    := filepath.Join(templateDir, "htmlOpen.tmpl")
  htmlCloseT   := filepath.Join(templateDir, "htmlClose.tmpl")
  formT        := filepath.Join(templateDir, "form.tmpl")
  t,err        := template.ParseFiles( layoutT, htmlOpenT, formT, htmlCloseT)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
  if err := t.ExecuteTemplate(w, "layout", formTitle); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
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
