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
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/person"
)

// type Crew is a struct that hold the crew data.
type Crew struct {
	ShipName  string
	Pilot     person.Person
	Navigator person.Person
	Medic     person.Person
	Engineers []person.Person
	Gunners   []person.Person
	Steward   person.Person
}

// type Ship is a struct that holds the ship data.
type Ship struct {
	ShipName   string
	HullSize   int
	DriveSize  int
	Passengers int
	Weapons    int
	Role       string
}

// buildCrew takes a Ship and a datadir and returns a Crew.
func buildCrew(ship Ship, datadir string) Crew {
	var crew Crew
	crew.ShipName = ship.ShipName
	var options = make(map[string]string)
	options["terms"] = "0"
	options["gender"] = ""
	options["datadir"] = datadir
	options["role"] = ship.Role
	options["job"] = "pilot"
	crew.Pilot = person.MakePerson(options)
	options["job"] = "navigator"
	crew.Navigator = person.MakePerson(options)
	if ship.DriveSize > 200 {
		ship.DriveSize = 200
	}
	for dT := ship.DriveSize; dT > 0; dT -= 35 {
		options["job"] = "engineer"
		crew.Engineers = append(crew.Engineers, person.MakePerson(options))
	}
	if ship.HullSize >= 200 {
		options["job"] = "medic"
		crew.Medic = person.MakePerson(options)
	}
	if ship.Passengers > 0 {
		options["job"] = "steward"
		crew.Steward = person.MakePerson(options)
	}
	if ship.Weapons > 0 {
		if ship.Weapons > 10 {
			ship.Weapons = 10
		}
		for g := 0; g < ship.Weapons; g++ {
			options["job"] = "gunner"
			crew.Gunners = append(crew.Gunners, person.MakePerson(options))
		}
	}
	return crew
}

// crewGen handles the http.Request and the http.ResponseWriter
func crewGen(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exedir := path.Dir(exe)
	datadir := path.Join(exedir, "data")
	webdir := path.Join(exedir, "web")
	// Needs more input sanitization.
	shipName := strings.Join(r.Form["shipName"], "")
	hullSize, _ := strconv.Atoi(strings.Join(r.Form["hullSize"], ""))
	driveSize, _ := strconv.Atoi(strings.Join(r.Form["driveSize"], ""))
	passengers, _ := strconv.Atoi(strings.Join(r.Form["passengers"], ""))
	weapons, _ := strconv.Atoi(strings.Join(r.Form["weapons"], ""))
	role := strings.Join(r.Form["role"], "")
	ship := Ship{ShipName: shipName,
		HullSize:   hullSize,
		DriveSize:  driveSize,
		Passengers: passengers,
		Weapons:    weapons,
		Role:       role,
	}
	crew := buildCrew(ship, datadir)
	layoutT := filepath.Join(webdir, "layout.tmpl")
	htmlOpenT := filepath.Join(webdir, "htmlOpen.tmpl")
	htmlCloseT := filepath.Join(webdir, "htmlClose.tmpl")
	formT := filepath.Join(webdir, "form.tmpl")
	crewT := filepath.Join(webdir, "crew.tmpl")
	personT := filepath.Join(webdir, "person.tmpl")
	t, err := template.ParseFiles(layoutT, htmlOpenT, htmlCloseT, formT, crewT, personT)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := t.ExecuteTemplate(w, "layout", crew); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default to use port %s", port)
	}
	http.HandleFunc("/", crewGen)
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
