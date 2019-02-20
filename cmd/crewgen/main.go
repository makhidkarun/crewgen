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

type Crew struct {
  ShipName    string
  Pilot       person.Person
  Navigator   person.Person
  Medic       person.Person
  Engineers   []person.Person
  Gunners     []person.Person
  Steward     person.Person
}

type Ship struct {
	ShipName    string
  HullSize    int
  DriveSize   int
  Passengers  int
  Weapons     int
}

func buildCrew(ship Ship) Crew {
  var crew Crew
  crew.ShipName = ship.ShipName
  var options = make(map[string]string)
  options["terms"] = "0"
  options["gender"] = ""
  options["db_name"] = "data/names.db"
  crew.Pilot = person.MakePerson(options)
  crew.Navigator = person.MakePerson(options)
  for dT := ship.DriveSize; dT > 0; dT -= 35 {
    crew.Engineers = append(crew.Engineers, person.MakePerson(options))
  }
  if ship.HullSize >= 200 {
    crew.Medic    = person.MakePerson(options)
  }
  if ship.Passengers > 0 {
    crew.Steward    = person.MakePerson(options)
  }
  if ship.Weapons > 0 {
    for g := 0; g < ship.Weapons; g++ {
      crew.Gunners   = append(crew.Gunners, person.MakePerson(options))
    }
  }

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
	ship := Ship{ShipName: shipName, 
    HullSize: hullSize,
    DriveSize: driveSize,
    Passengers: passengers,
    Weapons:  weapons,
  }
	crew := buildCrew(ship)
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
	if err := t.ExecuteTemplate(w, "layoutS", crew); err != nil {
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
