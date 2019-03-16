package person_test

import (
	"strings"
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

func TestMakePersonName(t *testing.T) {
	options := make(map[string]string)
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	if len(testP.Name) < 5 {
		t.Error(`MakePerson failed name`)
	}
}

func TestMakePersonNameTwoWordString(t *testing.T) {
	options := make(map[string]string)
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	nameS := strings.Split(testP.Name, " ")
  if len(nameS) != 2 {
		t.Error(`MakePerson failed two word name`)
	}
}

func TestCareerNavy(t *testing.T) {
	options := make(map[string]string)
	options["role"] = "Navy"
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	if testP.Career != "Navy" {
		t.Error(`MakePerson failed to specify Navy career`)
	}
}

func TestCareerMerchant(t *testing.T) {
	options := make(map[string]string)
	options["role"] = "MerchantMarine"
	options["db_name"] = "data/names.db"
	testP := person.MakePerson(options)
	if testP.Career != "Merchant" {
		t.Error(`MakePerson failed to specify Merchant career`)
	}
}

