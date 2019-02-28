package person_test

import (
	"testing"
	"fmt"
	"github.com/makhidkarun/crewgen/pkg/person"
)

func TestMakePerson(t *testing.T) {
	options := make(map[string]string)
	options["terms"] = "1"
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	fmt.Println(testP.Age)
	if testP.Age <= 21 || testP.Age >= 26 {
		t.Error(`MakePerson failed age`)
	}
}
