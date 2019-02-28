package person_test

import (
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

func TestMakePerson(t *testing.T) {
	options := make(map[string]string)
	options["terms"] = "1"
	options["db_name"] = "../../data/names.db"
	testP := person.MakePerson(options)
	if testP.Age == 22 {
		t.Error(`MakePerson failed age`)
	}
}
