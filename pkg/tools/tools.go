// Functions for Traveller and Cepheus Engine
// All copyrights to them that own's them.

package tools

import (
	"bufio"
	"os"

	"github.com/makhidkarun/crewgen/pkg/dice"
)

func ArrayFromFile(f *os.File) []string {
	items := make([]string, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if len(s) > 2 {
			items = append(items, s)
		}
	}
	return items
}

func StringInArray(val string, array []string) bool {
	for _, value := range array {
		if value == val {
			return true
		}
	}
	return false
}

func RandomStringFromArray(array []string) string {
	return array[dice.Random(0, len(array)-1)]
}
