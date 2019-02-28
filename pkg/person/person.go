// To create a basic 2d6 OGL character

// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike and the Google Go team,
// and Freenode#go-nuts.

package person

import (
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/tools"
)

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

func (p *Person) IncSkill(s string) {
	p.Skills[s] += 1
}

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

	character.Name = tools.GetName(character.Gender, db_name)
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
