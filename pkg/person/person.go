// Person provides the data structure and methods to create a basic
// 2d6 OGL character
//
// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike the Google Go team, Libera#go-nuts, and Slack/Gopher.

package person

import (
	crand "crypto/rand"
	"database/sql"
	"fmt"
	mbig "math/big"
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/dice"
	"github.com/makhidkarun/crewgen/pkg/tools"
	_ "github.com/mattn/go-sqlite3"
	//_ "modernc.org/sqlite"
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
	return 18 + (terms * 4) + rng(0, 3)
}

// newSkill takes a string and returns a string from a slice.
func newSkill(job string) (skill string) {
	switch job {
	case "engineer":
		skills := []string{"Engineering", "Engineering", "Engineering", "Mechanical", "Electronics"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	case "pilot":
		skills := []string{"Pilot", "Pilot", "Navigation", "Comms", "Sensors", "FleetTactics"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	case "navigator":
		skills := []string{"Navigation", "Navigation", "ShipsBoat", "Comms", "Sensors"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	case "medic":
		skills := []string{"Medical", "Medical", "Medical", "Diplomacy", "Science(Any)"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	case "gunner":
		skills := []string{"Gunnery", "Gunnery", "Brawling", "Mechanical", "Electronics"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	case "steward":
		skills := []string{"Steward", "Steward", "Diplomacy", "Carouse", "Medic"}
		n := rng(0, len(skills)-1)
		skill = skills[n]
	}
	return
}

// rng takes min and max ints and returns an int in the
// range min to max, inclusive.
func rng(min int, max int) int {
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

func setCareer(career ...string) (c string) {
	if career[0] == "Navy" {
		c = "Navy"
	} else if career[0] == "MerchantMarine" {
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
		if upp[stat] < 0 {
			upp[stat] = 0
		} else if upp[stat] > 15 {
			upp[stat] = 15
		}
	}
	return upp
}

func setGender() string {
	if dice.OneD6()%2 == 0 {
		return "F"
	} else {
		return "M"
	}
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

// numTerms sets the number of terms the character served.
func numTerms() (t int) {
	t = rng(1, 4)
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

// MakePerson takes a map of options and returns a Person.
// It is a basic factory.
func MakePerson(options map[string]string) Person {
	terms, _ := strconv.Atoi(options["terms"])
	gender := options["gender"]
	db_name := options["db_name"]
	career := options["role"]
	job := options["job"]

	speciesOptions := []string{"human", "human", "andorian", "human", "vulcan"}

	var character Person
	character.Skills = make(map[string]int)

	if terms <= 0 || terms >= 5 {
		character.Terms = numTerms()
	} else {
		character.Terms = terms
	}

	var genders []string = []string{"F", "M"}
	input_gender := strings.ToUpper(gender)
	if !tools.StringInArray(input_gender, genders) {
		character.Gender = setGender()
	} else {
		character.Gender = input_gender
	}

	character.Name = GetName(character.Gender, db_name)
	//character.UPP = tools.RollUPP()
	character.UPPs = formatUPP(character.UPP)
	character.Age = age(character.Terms)
	character.Career = setCareer(career)
	character.Species = setSpecies(speciesOptions)

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
	character.SkillString = character.SkillsToStr()
	character.Physical = writePhysical(character)

	return character
}
