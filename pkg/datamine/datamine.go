// datamine functions for Traveller and Cepheus Engine
// All copyrights to them that own's 'em.

package datamine

import (
	"fmt"
	"os"
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
