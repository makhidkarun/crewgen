// Person provides the data structure and methods to create a basic
// 2d6 OGL character
//
// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike the Google Go team, Libera#go-nuts, and Slack/Gopher.

package person

import (
	"fmt"
	"github.com/makhidkarun/crewgen/pkg/datamine"
	"github.com/makhidkarun/crewgen/pkg/dice"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
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
	Plot        string
	Mental      []string
}

// age sets the base age, assuming some time after leaving the service.
func age(options map[string]string, p Person) int {
	termMod, err := strconv.Atoi(options["termMod"])
	if err != nil {
		termMod = 4
	}
	return 18 + (p.Terms * termMod) + dice.Random(0, 3)
}

// newSkill takes an array and returns a string
func newSkill(job []string) string {
	return datamine.RandomStringFromArray(job)
}

// skillsToString returns a comma or newline separated single string.
func skillsToStr(options map[string]string, p Person) string {
	s := ""
	i := 1
	if len(p.Skills) == 0 {
		return s
	}
	skillList := make([]string, 0, len(p.Skills))
	for skill, _ := range p.Skills {
		skillList = append(skillList, skill)
	}
	sort.Strings(skillList)
	for _, key := range skillList {
		s += key + "-" + strconv.Itoa(p.Skills[key])
		if i < len(skillList) {
			if options["game"] == "brp" {
				s += "\n"
			} else {
				s += ", "
			}
			i++
		}
	}
	return s
}

// incSkill increases a skill by value of int
func incSkill(skills map[string]int, skill string, value int) map[string]int {
	skills[skill] += value
	return skills
}

// setCareer sets a random career if a valid string is not given.
func setCareer(options map[string]string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if _, err := os.Stat(cwd); err != nil {
		os.Exit(1)
	}
	var c string
	if _, err := os.Stat(options["careerFile"]); err != nil {
		fmt.Printf("%s does not exist error\n", options["careerFile"])
		os.Exit(1)
	}

	cOptions := datamine.CareerList(options["careerFile"])
	if !datamine.StringInArray(options["career"], cOptions) {
		c = datamine.RandomStringFromArray(cOptions)
	} else {
		c = options["career"]
	}
	return c
}

// getPlot selects a plot from the given file, and returns a string
func getPlot(options map[string]string) string {
	plotfile := path.Join(options["datadir"], "plots.txt")
	plot, err := datamine.RandomStringFromFile(plotfile)
	if err != nil {
		fmt.Println("Error getting a plot")
		os.Exit(1)
	}
	return plot
}

// formatUPP returns a string based on the game type.
func formatUPP(options map[string]string, p Person) string {
	var newUPP string
	if options["game"] == "brp" {
		statKeys := []string{"Str", "Con", "Siz", "Int", "Pow", "Dex"} //, "Edu"}
		for idx, val := range statKeys {
			newUPP += fmt.Sprintf("%s: %d  ", val, p.UPP[idx])
		}
	} else {
		for _, val := range p.UPP {
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
func setGender(options map[string]string) string {
	var genders []string = []string{"F", "M"}
	if len(options["gender"]) != 0 {
		if datamine.StringInArray(options["gender"], genders) {
			return options["gender"]
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
func rollUPP(options map[string]string) []int {
	stats := 6
	numDice := 2
	if options["game"] == "brp" {
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
func addSkills(options map[string]string, p Person) map[string]int {
	skills := make(map[string]int)
	careerList := datamine.CareerList(options["careerFile"])
	if !datamine.StringInArray(options["career"], careerList) {
		options["career"] = "other"
	}
	careerSkills := datamine.CareerSkills(options["careerFile"], options["career"])
	job := options["job"]
	if job == "" {
		job = datamine.DefaultJob(options["careerFile"], options["career"])
	} else {
		jobList := datamine.JobList(options["jobFile"])
		if !datamine.StringInArray(job, jobList) {
			job = "other"
		}
	}
	jobSkills := datamine.JobSkills(options["jobFile"], job)
	var skillList = make([]string, len(jobSkills)+len(careerSkills))
	copy(skillList, append(jobSkills, careerSkills[:]...))
	primarySkill := datamine.FirstStringInArray(skillList)
	incValue := 1
	if options["game"] == "brp" {
		incValue = 5
	}
	skills = incSkill(skills, primarySkill, incValue)
	for i := 0; i < p.Terms; i++ {
		skills = incSkill(skills, newSkill(skillList), incValue)
	}
	return skills
}

// MakePerson takes a map of options and returns a Person.
func MakePerson(options map[string]string) Person {
	var character Person

	terms, tErr := strconv.Atoi(options["terms"])
	if tErr != nil {
		character.Terms = numTerms()
	} else {
		character.Terms = terms
	}

	options["career"] = strings.ToLower(options["career"])
	options["job"] = strings.ToLower(options["job"])
	// This needs expansion, for each game type. Use in RandomStringFromArray
	speciesOptions := []string{"human"}

	if options["game"] == "brp" {
		options["termMod"] = "0"
	} else {
		options["termMod"] = "4"
	}
	character.Gender = setGender(options)
	character.UPP = rollUPP(options)
	character.Name = datamine.GetName(options)
	character.UPPs = formatUPP(options, character)
	character.Age = age(options, character)
	character.Career = setCareer(options)
	character.Species = setSpecies(speciesOptions)
	character.Skills = addSkills(options, character)
	character.SkillString = skillsToStr(options, character)
	character.Physical = writePhysical(character)
	character.Plot = getPlot(options)

	return character
}
