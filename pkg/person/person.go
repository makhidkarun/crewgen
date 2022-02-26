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
	"strconv"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/datamine"
	"github.com/makhidkarun/crewgen/pkg/dice"
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

// buildSkillList returns a map with a job (string) key and skills (array) data.
func buildSkillList() map[string][]string {
	skillList := make(map[string][]string)
	skillList["engineer"] = []string{"Engineering", "Engineering", "Mechanical", "Electronics"}
	skillList["pilot"] = []string{"Pilot", "Navigation", "Comms", "Sensors", "FleetTactics"}
	skillList["navigator"] = []string{"Navigation", "ShipsBoat", "Comms", "Sensors"}
	skillList["medic"] = []string{"Medical", "Medical", "Medical", "Diplomacy", "Science(Any)"}
	skillList["gunner"] = []string{"Gunnery", "Gunnery", "Brawling", "Mechanical", "Electronics"}
	skillList["steward"] = []string{"Steward", "Steward", "Diplomacy", "Carouse", "Medic"}
	skillList["infantry"] = []string{"GunCbt(CbtR)", "GunCbt(Any)", "HvyWpns(Any)", "Recon", "Drive(any)", "VaccSuit", "Brawling", "Gambling", "Mechanic", "Leader", "GunCbt(Any)"}
	skillList["commando"] = []string{"GunCbt(CbtR)", "GunCbt(Any)", "HvyWpns(Any)", "Demolition", "Survival", "Recon", "Battledress", "Leader", "Tactics", "Blade", "Instruction"}
	skillList["life"] = []string{"Drive(Any)", "Computer", "Admin", "Streetwise"}
	return skillList
}

// age sets the base age, assuming some time after leaving the service.
func age(terms int) int {
	return 18 + (terms * 4) + dice.Random(0, 3)
}

// newSkill takes an array and returns a string
func newSkill(job []string) string {
	return datamine.RandomStringFromArray(job)
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
func formatUPP(upp [6]int) string {
	var newUPP string
	for _, val := range upp {
		newUPP += fmt.Sprintf("%X", val)
	}
	return newUPP
}

// modifyUpp ensures all UPP numbers are between 2 and 15, decimal.
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
// Need to put this into the datamine.
func addSkills(job string, terms int) map[string]int {
	skills := make(map[string]int)
	skillList := buildSkillList()
	jobs := make([]string, 0, len(skillList))
	for j := range skillList {
		jobs = append(jobs, j)
	}
	if !datamine.StringInArray(job, jobs) {
		job = "life"
	}
	primarySkill := datamine.FirstStringInArray(skillList[job])

	skills = incSkill(skills, primarySkill)
	for i := 0; i < terms; i++ {
		skills = incSkill(skills, newSkill(skillList[job]))
	}
	return skills
}

// MakePerson takes a map of options and returns a Person.
// It is a basic factory.
func MakePerson(options map[string]string) Person {
	terms, _ := strconv.Atoi(options["terms"])
	input_gender := options["gender"]
	datadir := options["datadir"]
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
	character.Name = datamine.GetName(character.Gender, datadir)
	character.UPP = rollUPP()
	character.UPPs = formatUPP(character.UPP)
	character.Age = age(character.Terms)
	character.Career = setCareer(career, datadir)
	character.Species = setSpecies(speciesOptions)
	character.Skills = addSkills(job, character.Terms)
	character.SkillString = skillsToStr(character.Skills)
	character.Physical = writePhysical(character)

	return character
}
