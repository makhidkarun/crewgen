// Person provides the data structure and methods to create a basic
// 2d6 OGL character
//
// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike the Google Go team, Libera#go-nuts, and Slack/Gopher.

package person

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/datamine"
	"github.com/makhidkarun/crewgen/pkg/dice"
)

// Person holds data. Most fields are exported.
type Person struct {
	Name        string
	UPP         []int
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
func age(terms int, termMod int) int {
	return 18 + (terms * termMod) + dice.Random(0, 3)
}

// newSkill takes an array and returns a string
func newSkill(job []string) string {
	return datamine.RandomStringFromArray(job)
}

// skillsToString returns a comma or newline separated single string.
func skillsToStr(skills map[string]int, game string) string {
	s := ""
	i := 1
	if len(skills) == 0 {
		return s
	}
	skillList := make([]string, 0, len(skills))
	for skill, _ := range skills {
		skillList = append(skillList, skill)
	}
	sort.Strings(skillList)
	for _, key := range skillList {
		s += key + "-" + strconv.Itoa(skills[key])
		if i < len(skillList) {
			if game == "brp" {
				s += "\n"
			} else {
				s += ", "
			}
			i++
		}
	}
	return s
}

// incSkill increases a skill by 1
// probably needs a variable?
func incSkill(skills map[string]int, skill string, value int) map[string]int {
	skills[skill] += value
	return skills
}

// setCareer sets a random career if a valid string is not given.
func setCareer(career string, datadir string) (c string) {
	datafile := path.Join(datadir, "careers.txt")
	cOptions := datamine.CareerList(datafile)
	if !datamine.StringInArray(career, cOptions) {
		c = datamine.RandomStringFromArray(cOptions)
	} else {
		c = career
	}
	return
}

// formatUPP returns a string of alphanumeric Hex.
func formatUPP(upp []int, game string) string {
	var newUPP string
	if game == "brp" {
		statKeys := []string{"Str", "Con", "Siz", "Int", "Pow", "Dex"} //, "Edu"}
		for idx, val := range statKeys {
			newUPP += fmt.Sprintf("%s: %d  ", val, upp[idx])
		}
	} else {
		for _, val := range upp {
			newUPP += fmt.Sprintf("%X", val)
		}
	}
	return newUPP
}

// modifyUpp ensures all UPP numbers are between 2 and 15, decimal.
func modifyUpp(upp []int, stat int, mod int) []int {
	// Requires the UPP [6]int, stat index [0-5], and modifier
	if stat < 0 || stat >= len(upp) {
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

// setGender allows choosing a gender from M, F, or random assignment.
func setGender(input ...string) string {
	var genders []string = []string{"F", "M"}
	if len(input) != 0 {
		test_gender := strings.ToUpper(input[0])
		if datamine.StringInArray(test_gender, genders) {
			return test_gender
		}
	}
	return datamine.RandomStringFromArray(genders)
}

// numTerms sets the number of terms the character served.
func numTerms() (t int) {
	t = dice.Random(1, 5)
	return
}

// rollUPP starts the basic 2d6 rolls.
func rollUPP(game string) []int {
	stats := 6
	numDice := 2
	if game == "brp" {
		stats = 8
		numDice = 3
	}
	upp := make([]int, 0, stats)
	for i := 0; i < stats; i++ {
		roll := 0
		for j := 0; j < numDice; j++ {
			roll += dice.OneD6()
		}
		upp = append(upp, roll)
	}
	return upp
}

// setSpecies takes a list of options and assigns one randomly to the character.
func setSpecies(l []string) (species string) {
	species = datamine.RandomStringFromArray(l)
	return
}

// writePhysical takes a Person and returns a physical description string.
func writePhysical(c Person) string {
	hOptions := []string{"short", "medium height", "tall"}
	wOptions := []string{"thin", "medium build", "heavy set"}
	height := datamine.RandomStringFromArray(hOptions)
	weight := datamine.RandomStringFromArray(wOptions)
	gen := map[string]string{"F": "female", "M": "male"}
	physical := fmt.Sprintf("%s is a %s, %s %s %s",
		c.Name, height, weight, c.Species, gen[c.Gender])
	return physical
}

// addSkills returns a map of skill (string) and level (int).
//  It auto assigns the primary skill for the job given.
func addSkills(job string, career string, terms int, datadir string, game string) map[string]int {
	skills := make(map[string]int)
	careerFile := path.Join(datadir, "careers.txt")
	careerList := datamine.CareerList(careerFile)
	if !datamine.StringInArray(career, careerList) {
		career = "Other"
	}
	careerSkills := datamine.CareerSkills(careerFile, career)

	jobFile := path.Join(datadir, "jobs.txt")
	if job == "" {
		job = datamine.DefaultJob(careerFile, career)
	} else {
		jobList := datamine.JobList(jobFile)
		if !datamine.StringInArray(job, jobList) {
			job = "other"
		}
	}
	jobSkills := datamine.JobSkills(jobFile, job)
	var skillList = make([]string, len(jobSkills)+len(careerSkills))
	copy(skillList, append(jobSkills, careerSkills[:]...))
	primarySkill := datamine.FirstStringInArray(skillList)

	value := 1
	if game == "brp" {
		value = 5
	}
	skills = incSkill(skills, primarySkill, value)
	for i := 0; i < terms; i++ {
		skills = incSkill(skills, newSkill(skillList), value)
	}
	return skills
}

// MakePerson takes a map of options and returns a Person.
// It is a basic factory.
func MakePerson(options map[string]string) Person {
	var character Person

	terms, tErr := strconv.Atoi(options["terms"])
	if tErr != nil {
		character.Terms = numTerms()
	} else {
		character.Terms = terms
	}

	input_gender := options["gender"]
	datadir := options["datadir"]
	career := strings.ToLower(options["career"])
	job := strings.ToLower(options["job"])
	game := options["game"]
	lastName := options["lastName"]

	speciesOptions := []string{"human"}
	termMod := 4
	if game == "brp" {
		termMod = 0
	}
	character.Gender = setGender(input_gender)
	character.Name = datamine.GetName(character.Gender, datadir, lastName)
	character.UPP = rollUPP(game)
	character.UPPs = formatUPP(character.UPP, game)
	character.Age = age(character.Terms, termMod)
	character.Career = setCareer(career, datadir)
	character.Species = setSpecies(speciesOptions)
	character.Skills = addSkills(job, character.Career, character.Terms, datadir, game)
	character.SkillString = skillsToStr(character.Skills, game)
	character.Physical = writePhysical(character)

	return character
}
