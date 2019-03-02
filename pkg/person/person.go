// Person provides the data structure and methods to create a basic 
// 2d6 OGL character
//
// With acknowledgement of Marc Miller of FFE, 
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike the Google Go team, Freenode#go-nuts, and Slack/Gopher.

package person

import (
	"database/sql"
	"fmt"
	crand "crypto/rand"
	mbig "math/big"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/makhidkarun/crewgen/pkg/tools"
)

// Person holds data. Most fields are exported.
type Person struct {
	Name   string
	UPP    [6]int
	UPPs   string
	Gender string
	PSR    int
	Age    int
	Terms  int
	Career string
	Skills map[string]int
	S      string
}

// newSkill takes a string and returns a string from a slice.
func newSkill( job string ) (skill string) {
	switch job {
		case "engineer":
			skills := []string{ "Engineering", "Engineering", "Engineering", "Mechanical", "Electronics", }
			n := RNG(1, len(skills) - 1 )
			skill = skills[n]
		case "pilot":
			skills := []string{ "Pilot", "Pilot", "Navigation", "Comms", "Sensors", "FleetTactics" }
			n := RNG(1, len(skills) - 1 )
			skill = skills[n]
		case "navigator":
			skills := []string{ "Navigation", "Navigation", "ShipsBoat", "Comms", "Sensors" }
			n := RNG(1, len(skills) - 1 )
			skill = skills[n]
		case "medic":
			skills := []string{ "Medical", "Medical", "Medical", "Diplomacy", "Science(Any)", }
			n := RNG(1, len(skills) - 1 )
			skill = skills[n]
		case "gunner":
			skills := []string{ "Gunnery", "Gunnery", "Brawling", "Mechanical", "Electronics", }
			//n := RNG(1, len(skills) - 1 )
			n := RNG(0,4)
			skill = skills[n]
		case "steward":
			skills := []string{ "Steward", "Steward", "Diplomacy", "Carouse", "Medic"}
			n := RNG(1, len(skills) - 1 )
			skill = skills[n]
	}
	return
}

// RNG takes min and max ints and returns an int in the 
// range min to max, inclusive.
func RNG(min int, max int) int {
	spread := max - min + 1
	spread64 := int64(spread)
	num, _ := crand.Int(crand.Reader, mbig.NewInt(spread64))
	roll := int(num.Int64()) + min
	return roll
}

// SkillsToString returns a comma separate single string.
//   Skill-1,Skill-3
func (p *Person) SkillsToStr() (s string) {
	i := 1
	for k, v := range p.Skills {
		s += k + "-" + strconv.Itoa(v)
		if i < len(p.Skills) {
			i++
			s += ", "
		}
	}
	return
}

// IncSkill takes a string of skill as map key and increments the value by 1.
func (p *Person) IncSkill(s string) {
	p.Skills[s] += 1
}

// GetName takes a string of "F" or "M" and a string of a database location.
// Returns a string of "FirstName LastName".
// Uses a SQLite3 database, "database/sql", and "github.com/mattn/go-sqlite3".
func GetName(gender string, db_name string) string {
	// Note that the names.db file must be where the command is run
	// from.

	var lname string
	var fname string
	var fresult *sql.Rows

	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// First Name
	if gender == "M" {
		fresult, err = db.Query("SELECT name FROM humaniti_male_first ORDER BY RANDOM() LIMIT 1")
	} else {
		fresult, err = db.Query("SELECT name FROM humaniti_female_first ORDER BY RANDOM() LIMIT 1")
	}
	if err != nil {
		fmt.Println(err)
	}
	for fresult.Next() {
		err = fresult.Scan(&fname)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Last Name
	result, err := db.Query("SELECT name FROM humaniti_last ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err = result.Scan(&lname)
		if err != nil {
			fmt.Println(err)
		}
	}

	name := fname + " " + lname
	return name
}

func numTerms() (t int) {
	t = RNG(1,4)
	return
}

// MakePerson takes a map of options and returns a Person.
// It is a basic factory.
func MakePerson(options map[string]string) Person {
	terms, _ := strconv.Atoi(options["terms"])
	gender := options["gender"]
	db_name := options["db_name"]
	career := options["role"]
	job := options["job"]

	var character Person
	character.Skills = make(map[string]int)

	if ( terms <= 0 || terms >=5 ) {
		character.Terms = numTerms()
	} else {
		character.Terms = terms
	}

	var genders []string = []string{"F", "M"}
	input_gender := strings.ToUpper(gender)
	if !tools.StringInArray(input_gender, genders) {
		character.Gender = tools.Gender()
	} else {
		character.Gender = input_gender
	}

	character.Name = GetName(character.Gender, db_name)
	character.UPP = tools.RollUPP()
	character.UPPs = tools.FormatUPP(character.UPP)
	character.Age = tools.Age(character.Terms)
	character.Career = tools.Career(career)

	var primarySkill string
	var nS string
	
	switch job {
	case "pilot":
		primarySkill = "Pilot"
	case "navigator":
		primarySkill = "Navigator"
	case "engineer":
		primarySkill = "Engineering"
	case "gunner":
		primarySkill = "Gunnery"
	case "medic":
		primarySkill = "Medical"
	case "steward":
		primarySkill = "Steward"
	}
	character.IncSkill(primarySkill)
	for i := 0; i < character.Terms; i++ {
		nS = newSkill(job)
		character.IncSkill(nS)
	}
	character.S = character.SkillsToStr()
	return character
}
