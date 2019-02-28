package person_test

import (
	"reflect"
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

func TestMakePerson(t *testing.T) {
	options := make(map[string]string)
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
  testS := "test string"
	if reflect.TypeOf(testP.Name) != reflect.TypeOf(testS) {
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

