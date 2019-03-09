// crewgen provides a list of characters to crew a starship.
// Output is based on ship hull size, drive tonnage, passengers, and weapons aboard.
// Logic is similar to 2d6 OGL and other Sci-Fi role-playing games.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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
	Role        string
}

func buildCrew(ship Ship) Crew {
	var crew Crew
	crew.ShipName = ship.ShipName
	var options = make(map[string]string)
	options["terms"] = "0"
	options["gender"] = ""
	options["db_name"] = "data/names.db"
	options["role"] = ship.Role
	options["job"]  = "pilot"
	crew.Pilot = person.MakePerson(options)
	options["job"]  = "navigator"
	crew.Navigator = person.MakePerson(options)
	if ship.DriveSize > 200 {
		ship.DriveSize = 200
	}
	for dT := ship.DriveSize; dT > 0; dT -= 35 {
		options["job"]  = "engineer"
		crew.Engineers = append(crew.Engineers, person.MakePerson(options))
	}
	if ship.HullSize >= 200 {
		options["job"]  = "medic"
		crew.Medic    = person.MakePerson(options)
	}
	if ship.Passengers > 0 {
		options["job"]  = "steward"
		crew.Steward    = person.MakePerson(options)
	}
	if ship.Weapons > 0 {
		if ship.Weapons > 10 {
			ship.Weapons = 10
		}
		for g := 0; g < ship.Weapons; g++ {
			options["job"]  = "gunner"
			crew.Gunners   = append(crew.Gunners, person.MakePerson(options))
		}
	}
	return crew
}

func showCrew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Needs more input sanitization.
	shipName      := strings.Join(r.Form["shipName"], "")
	hullSize,_    := strconv.Atoi(strings.Join(r.Form["hullSize"], ""))
	driveSize,_   := strconv.Atoi(strings.Join(r.Form["driveSize"], "")) 
	passengers,_  := strconv.Atoi(strings.Join(r.Form["passengers"], "")) 
	weapons,_     := strconv.Atoi(strings.Join(r.Form["weapons"], "")) 
	role          := strings.Join(r.Form["role"], "")
	ship := Ship{ShipName: shipName, 
		HullSize: hullSize,
		DriveSize: driveSize,
		Passengers: passengers,
		Weapons:  weapons,
		Role:     role,
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
	//server := http.Server{
	//	Addr: "127.0.0.1:8080",
	//}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default to use port %s", port)
	}
	http.HandleFunc("/recruit", recruitCrew)
	http.HandleFunc("/show", showCrew)
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

