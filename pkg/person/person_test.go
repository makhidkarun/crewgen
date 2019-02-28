package person_test

import (
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

func TestMakePerson(t *testing.T) {
	options := make(map[string]string)
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	var tP interface{} = testP.Name
  if _, ok := tP.(string); ! ok {
    t.Error(`MakePerson failed by name test`)
  }	
}
 

func TestMakePersonAge(t *testing.T) {
	options := make(map[string]string)
	options["terms"] = "1"
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	if testP.Age <= 21 || testP.Age >= 26 {
		t.Error(`MakePerson failed age`)
	}
}

