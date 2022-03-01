// datamine functions for Traveller and Cepheus Engine
// All copyrights to them that own's 'em.

package datamine

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/makhidkarun/crewgen/pkg/dice"
)

//var careerList []string
//var jobList []string

func LineToList(line string, sep string) []string {
	var data []string
	for _, item := range strings.Split(line, sep) {
		data = append(data, strings.Trim(item, " "))
	}
	return data
}

func HeadersFromList(data []string, sep string) []string {
	// assumes blank lines and comments have been filtered out.
	var headers []string
	for _, line := range data {
		datum := strings.Split(line, sep)[0]
		headers = append(headers, strings.Trim(datum, " "))
	}
	return headers
}

func DataFromListLine(data []string, key string, sep string, index int) string {
	var datum string
	for _, line := range data {
		if strings.HasPrefix(line, key) {
			if sep != "" && index >= 0 {
				datum = strings.Split(line, sep)[index]
			} else {
				datum = line
			}
			break
		}
	}
	return datum
}

// whine prints the error message but keeps on going.
func whine(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ArrayFromFile takes a string filename and returns an array of strings,
//  one array item per file line.
func ArrayFromFile(filename string) []string {
	var items []string
	file, err := os.ReadFile(filename)
	whine(err)
	unwashed_items := strings.Split(string(file), "\n")
	for _, item := range unwashed_items {
		if !strings.HasPrefix(item, "#") {
			item = strings.Trim(item, " \n\t")
			if len(item) > 0 {
				items = append(items, item)
			}
		}
	}
	return items
}

// StringInArray takes a string to search for, and an array to search in.
// Returns true if the exact match of the string is in the array.
func StringInArray(val string, array []string) bool {
	for _, value := range array {
		if value == val {
			return true
		}
	}
	return false
}

// RandomStringFromArray takes an array of string and returns a random string.
func RandomStringFromArray(array []string) string {
	return array[dice.Random(0, len(array)-1)]
}

// FirstStringInArray takes an array and returns the first string.
func FirstStringInArray(array []string) string {
	return array[0]
}

// GetName takes a gender of string and returns a string of first and last names.
func GetName(gender string, datadir string) string {
	var first_name_file string
	if gender == "F" {
		first_name_file = path.Join(datadir, "human_female_first.txt")
	} else {
		first_name_file = path.Join(datadir, "human_male_first.txt")
	}

	first_name_list := ArrayFromFile(first_name_file)
	first_name := RandomStringFromArray(first_name_list)

	last_name_file := path.Join(datadir, "human_last.txt")
	last_name_list := ArrayFromFile(last_name_file)
	last_name := RandomStringFromArray(last_name_list)
	name := fmt.Sprintf("%s %s", first_name, last_name)
	return name
}

// CareerList provides the career options based on datafiles.
func CareerList(careerFile string) []string {
	sep := ":"
	careerData := ArrayFromFile(careerFile)
	careerList := HeadersFromList(careerData, sep)
	return careerList
}

// CareerSkills provides the skill options for specific careers.
// Format is "career:Skill, Skill, Skill, Skill
func CareerSkills(careerFile, career string) []string {
	careerSep := ":"
	skillSep := ","
	careerData := ArrayFromFile(careerFile)
	careerLine := DataFromListLine(careerData, career, careerSep, 2)
	careerSkills := LineToList(careerLine, skillSep)
	return careerSkills
}

// DefaultJob provides the default job, if the career exists.
// Otherwise "other".
func DefaultJob(careerFile, career string) string {
	careerSep := ":"
	careerData := ArrayFromFile(careerFile)
	defaultJob := DataFromListLine(careerData, career, careerSep, 1)
	return defaultJob
}

// JobList provides the job options based on datafiles.
func JobList(jobFile string) []string {
	sep := ":"
	jobData := ArrayFromFile(jobFile)
	jobList := HeadersFromList(jobData, sep)
	return jobList
}

// JobSkills provides the skill options for specific jobs.
// Format is "job:Skill, Skill, Skill, Skill
func JobSkills(jobFile, job string) []string {
	jobSep := ":"
	skillSep := ","
	jobData := ArrayFromFile(jobFile)
	jobLine := DataFromListLine(jobData, job, jobSep, 1)
	jobSkills := LineToList(jobLine, skillSep)
	return jobSkills
}
