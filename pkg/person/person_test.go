package person_test

import (
	"os"
	"strings"
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

var options map[string]string

func TestMain(m *testing.M) {
	options = make(map[string]string)
	options["db_name"] = "data/names.db"
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestMakePerson(t *testing.T) {
	testP := person.MakePerson(options)
	var tP interface{} = testP.Name
	if _, ok := tP.(string); !ok {
		t.Error(`MakePerson failed by name test`)
	}
}

func TestMakePersonAge(t *testing.T) {
	options["terms"] = "1"
	testP := person.MakePerson(options)
	if testP.Age <= 21 || testP.Age >= 26 {
		t.Error(`MakePerson failed age`)
	}
}

func TestMakePersonName(t *testing.T) {
	testP := person.MakePerson(options)
	if len(testP.Name) < 5 {
		t.Error(`MakePerson failed name`)
	}
}

func TestMakePersonNameTwoWordString(t *testing.T) {
	testP := person.MakePerson(options)
	nameS := strings.Split(testP.Name, " ")
	if len(nameS) != 2 {
		t.Error(`MakePerson failed two word name`)
	}
}

func TestCareerNavy(t *testing.T) {
	options["role"] = "Navy"
	testP := person.MakePerson(options)
	if testP.Career != "Navy" {
		t.Error(`MakePerson failed to specify Navy career`)
	}
}

func TestCareerMerchant(t *testing.T) {
	options["role"] = "MerchantMarine"
	testP := person.MakePerson(options)
	if testP.Career != "Merchant" {
		t.Error(`MakePerson failed to specify Merchant career`)
	}
}

func TestSpecies(t *testing.T) {
	testP := person.MakePerson(options)
	speciesOptions := map[string]bool{"human": true, "andorian": true, "vulcan": true}
	if speciesOptions[testP.Species] != true {
		t.Error(`MakePerson failed to specify valid species`)
	}
}

func TestPhysical(t *testing.T) {
	testP := person.MakePerson(options)
	if len(testP.Physical) < 10 {
		t.Error(`MakePerson failed to specify valid physical`)
	}
}
