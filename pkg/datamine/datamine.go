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

func whine(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// ArrayFromFile takes a string filename and returns an array of strings,
//  one array item per file line.
func ArrayFromFile(filename string) []string {
	file, err := os.ReadFile(filename)
	whine(err)
	items := strings.Split(string(file), "\n")
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
