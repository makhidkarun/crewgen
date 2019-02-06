package main

import (
  "fmt"
	//"io/ioutil"
  "log"
	"net/http"

	"github.com/makhidkarun/crewgen/pkg/person"
)

func showCrew(w http.ResponseWriter, r *http.Request) {
	p := person.Person{Name: "Al Lefron"}
  fmt.Fprintf(w, "Hey, I'm %s.", p.Name)
}

func main() {
  http.HandleFunc("/", showCrew) 
  log.Fatal(http.ListenAndServe(":8080", nil))
}
