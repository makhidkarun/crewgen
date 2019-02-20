// crewgen provides a list of characters to crew a starship.
// Output is based on ship hull size, drive tonnage, passengers, and weapons aboard.
// Logic is similar to 2d6 OGL and other Sci-Fi role-playing games.

package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
  "strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/person"
)

var templateDir string = "web"

type Crew []person.Person
type Ship struct {
	ShipName    string
	Crew        []person.Person
  HullSize    int
  DriveSize   int
  Passengers  int
  Weapons     int
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

func showCrew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
  // Needs more input sanitization.
	shipName := strings.Join(r.Form["shipName"], "")
  hullSize,_  := strconv.Atoi(strings.Join(r.Form["hullSize"], ""))
  driveSize,_  := strconv.Atoi(strings.Join(r.Form["driveSize"], "")) 
  passengers,_  := strconv.Atoi(strings.Join(r.Form["passengers"], "")) 
  weapons,_  := strconv.Atoi(strings.Join(r.Form["weapons"], "")) 
	crew := buildCrew()
	ship := Ship{ShipName: shipName, 
    Crew: crew, 
    HullSize: hullSize,
    DriveSize: driveSize,
    Passengers: passengers,
    Weapons:  weapons,
  }
	layoutST := filepath.Join(templateDir, "layoutS.tmpl")
	htmlOpenT := filepath.Join(templateDir, "htmlOpen.tmpl")
	htmlCloseT := filepath.Join(templateDir, "htmlClose.tmpl")
	formT := filepath.Join(templateDir, "form.tmpl")
	crewT := filepath.Join(templateDir, "crew.tmpl")
	personT := filepath.Join(templateDir, "person.tmpl")
	t, err := template.ParseFiles(layoutST, htmlOpenT, htmlCloseT, formT, crewT, personT)
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

func recruitCrew(w http.ResponseWriter, r *http.Request) {
	formTitle := "Recruiting Your Crew"
	layoutT := filepath.Join(templateDir, "layout.tmpl")
	htmlOpenT := filepath.Join(templateDir, "htmlOpen.tmpl")
	htmlCloseT := filepath.Join(templateDir, "htmlClose.tmpl")
	formT := filepath.Join(templateDir, "form.tmpl")
	t, err := template.ParseFiles(layoutT, htmlOpenT, formT, htmlCloseT)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
	if err := t.ExecuteTemplate(w, "layout", formTitle); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/recruit", recruitCrew)
	http.HandleFunc("/show", showCrew)
	log.Println("Starting server.")
	log.Fatal(server.ListenAndServe())
}
