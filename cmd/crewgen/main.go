package main

import (
  "fmt"
	//"io/ioutil"
  "html/template"
  "log"
	"net/http"

	"github.com/makhidkarun/crewgen/pkg/person"
)

type Page struct {
  Title string
  Body []byte
}

var templateDir string = "../../web/"

func loadPage(title string) (*Page, error) {
	p := person.Person{Name: "Al Lefron"}
  body := p.Name
  return &Page{Title: title, Body: []byte(body)}, nil
}
 
func recruitCrew(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Recruiting </h1>"           +
    "<form action=\"showCrew\" method=\"POST\">"   +
    "<input type=\"text\" name=\"shipName\"</br>" +
    "<input type=\"submit\" value=\"Recruit!\">"  +
    "</form>")
}
    
func showCrew(w http.ResponseWriter, r *http.Request) {
  title   := "The Crew"
  page, _ := loadPage(title)
  t, _    := template.ParseFiles( templateDir + "crew.html")
  t.Execute(w, page)
}

func main() {
  http.HandleFunc("/recruit", recruitCrew)
  http.HandleFunc("/showCrew", showCrew) 
  log.Fatal(http.ListenAndServe(":8080", nil))
}
