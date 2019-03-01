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

	// Need to figure out pre-adult characters
	if terms <= 0 {
		character.Terms = tools.NumTerms()
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
	character.S = character.SkillsToStr()
	return character
}
