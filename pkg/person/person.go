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
	Gender string
	PSR    int
	Age    int
	Terms  int
	Career string
}

func MakePerson(options map[string]string) Person {
	terms, _ := strconv.Atoi(options["terms"])
	gender := options["gender"]
	db_name := options["db_name"]

	var character Person

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
	character.Age = tools.Age(character.Terms)

	return character
}
