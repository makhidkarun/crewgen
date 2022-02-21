// Person provides the data structure and methods to create a basic
// 2d6 OGL character
//
// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike the Google Go team, Libera#go-nuts, and Slack/Gopher.

package person

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/dice"
	"github.com/makhidkarun/crewgen/pkg/tools"
	_ "github.com/mattn/go-sqlite3"
)

// Person holds data. Most fields are exported.
type Person struct {
	Name        string
	UPP         [6]int
	UPPs        string
	Gender      string
	PSR         int
	Age         int
	Terms       int
	Career      string
	Skills      map[string]int
	SkillString string
	Species     string
	Physical    string
}

// age sets the base age, assuming some time after leaving the service.
func age(terms int) int {
	return 18 + (terms * 4) + dice.Random(0, 3)
}

// newSkill takes a string and returns a string from a slice.
func newSkill(job string) (skill string) {
	switch job {
	case "engineer":
		skills := []string{"Engineering", "Engineering", "Engineering", "Mechanical", "Electronics"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	case "pilot":
		skills := []string{"Pilot", "Pilot", "Navigation", "Comms", "Sensors", "FleetTactics"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	case "navigator":
		skills := []string{"Navigation", "Navigation", "ShipsBoat", "Comms", "Sensors"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	case "medic":
		skills := []string{"Medical", "Medical", "Medical", "Diplomacy", "Science(Any)"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	case "gunner":
		skills := []string{"Gunnery", "Gunnery", "Brawling", "Mechanical", "Electronics"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	case "steward":
		skills := []string{"Steward", "Steward", "Diplomacy", "Carouse", "Medic"}
		n := dice.Random(0, len(skills)-1)
		skill = skills[n]
	}
	return
}

// skillsToString returns a comma separate single string.
//   Skill-1,Skill-3
func skillsToStr(skills map[string]int) string {
	s := ""
	i := 1
	if len(skills) == 0 {
		return s
	}
	for k, v := range skills {
		s += k + "-" + strconv.Itoa(v)
		if i < len(skills) {
			i++
			s += ", "
		}
	}
	return s
}

// incSkill increases a skill by 1
// probably needs a variable?
func incSkill(skills map[string]int, skill string) map[string]int {
	skills[skill] += 1
	return skills
}

func setCareer(career ...string) (c string) {
	if career[0] == "Navy" {
		c = "Navy"
	} else if career[0] == "Merchant" {
		c = "Merchant"
	} else {
		cOptions := []string{"Navy", "Merchant"}
		c = tools.RandomStringFromArray(cOptions)
	}
	return
}

func formatUPP(upp [6]int) string {
	var newUPP string
	for _, val := range upp {
		newUPP += fmt.Sprintf("%X", val)
	}
	return newUPP
}

func modifyUpp(upp [6]int, stat int, mod int) [6]int {
	// Requires the UPP [6]int, stat index [0-5], and modifier
	if stat < 0 || stat > 5 {
		return upp
	} else {
		upp[stat] += mod
		if upp[stat] < 2 {
			upp[stat] = 2
		} else if upp[stat] > 15 {
			upp[stat] = 15
		}
	}
	return upp
}

// stringInArray returns true if an exact string match is in an array.
// Returns false otherwise.
func stringInArray(val string, array []string) bool {
	for _, value := range array {
		if value == val {
			return true
		}
	}
	return false
}

// setGender allows choosing a gender from M, F, or random assignment.
// This could better be done with datamine.RandomStringFromArray()
func setGender(input ...string) string {
	var genders []string = []string{"F", "M"}
	if len(input) != 0 {
		test_gender := strings.ToUpper(input[0])
		if stringInArray(test_gender, genders) {
			return test_gender
		}
	}
	if dice.OneD6()%2 == 0 {
		return "F"
	} else {
		return "M"
	}
}

// getName takes a string of "F" or "M" and a string of a database location.
// Returns a string of "FirstName LastName".
func getName(gender string, db_name string) string {
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

// numTerms sets the number of terms the character served.
func numTerms() (t int) {
	t = dice.Random(1, 4)
	return
}

// rollUPP starts the basic 2d6 rolls.
func rollUPP() [6]int {
	var upp [6]int
	for i := 0; i < 6; i++ {
		upp[i] = dice.TwoD6()
	}
	return upp
}

// setSpecies takes a list of options and assigns one randomly to the character.
func setSpecies(l []string) (species string) {
	species = tools.RandomStringFromArray(l)
	return
}

// writePhysical takes a Person and returns a physical description string.
func writePhysical(c Person) string {
	hOptions := []string{"short", "medium height", "tall"}
	wOptions := []string{"thin", "medium build", "heavy set"}
	height := tools.RandomStringFromArray(hOptions)
	weight := tools.RandomStringFromArray(wOptions)
	gen := map[string]string{"F": "female", "M": "male"}
	physical := fmt.Sprintf("%s is a %s, %s %s %s",
		c.Name, height, weight, c.Species, gen[c.Gender])
	return physical
}

func addSkills(job string, terms int) map[string]int {
	skills := make(map[string]int)

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
	default:
		return skills
	}

	skills = incSkill(skills, primarySkill)
	for i := 0; i < terms; i++ {
		nS = newSkill(job)
		skills = incSkill(skills, nS)
	}
	return skills
}

// MakePerson takes a map of options and returns a Person.
// It is a basic factory.
func MakePerson(options map[string]string) Person {
	terms, _ := strconv.Atoi(options["terms"])
	input_gender := options["gender"]
	db_name := options["db_name"]
	career := options["career"]
	job := options["job"]

	speciesOptions := []string{"human"}
	var character Person

	if terms <= 0 || terms > 5 {
		character.Terms = numTerms()
	} else {
		character.Terms = terms
	}

	character.Gender = setGender(input_gender)
	character.Name = getName(character.Gender, db_name)
	character.UPP = rollUPP()
	character.UPPs = formatUPP(character.UPP)
	character.Age = age(character.Terms)
	character.Career = setCareer(career)
	character.Species = setSpecies(speciesOptions)
	character.Skills = addSkills(job, character.Terms)
	character.SkillString = skillsToStr(character.Skills)
	character.Physical = writePhysical(character)

	return character
}
