package person_test

import (
	"os"
	//"path"
	"strings"
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

var options map[string]string

func TestMain(m *testing.M) {
	options = make(map[string]string)
	//options["db_name"] = "data/names.db"
	/*exe, err := os.Executable()
	if err != nil {
		fmt.Println(`exe failed`)
	}
	exedir := path.Dir(exe)
	//options["datadir"] = path.Join(exedir, "data")
	*/
	datadir := "/home/leam/lang/git/makhidkarun/crewgen/cmd/teamgen/data"
	options["terms"] = "1"
	options["career"] = "Navy"
	options["datadir"] = datadir
	//options["job"] = "pilot"
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
	options["career"] = "Navy"
	testP := person.MakePerson(options)
	if testP.Career != "Navy" {
		t.Error(`MakePerson failed to specify Navy career`)
	}
}

func TestCareerMerchant(t *testing.T) {
	options["career"] = "Merchant"
	testP := person.MakePerson(options)
	if testP.Career != "Merchant" {
		t.Error(`MakePerson failed to specify Merchant career`)
	}
}

func TestSpecies(t *testing.T) {
	testP := person.MakePerson(options)
	speciesOptions := map[string]bool{"human": true}
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

func TestSkills(t *testing.T) {
	options["terms"] = "4"
	options["job"] = "pilot"
	testP := person.MakePerson(options)
	if len(testP.SkillString) < 8 {
		t.Error(`MakePerson failed to specify a long skillstring`)
	}
	if strings.Index(testP.SkillString, ", ") < 4 {
		t.Errorf("MakePerson does not have a comma and space in skillstring.")
	}
}

func TestNoSkills(t *testing.T) {
	options["terms"] = "4"
	options["job"] = ""
	testP := person.MakePerson(options)
	if len(testP.SkillString) > 0 {
		t.Errorf("MakePerson didn't do a blank skillstring %q\n", testP.SkillString)
	}
}
func TestUPP(t *testing.T) {
	testP := person.MakePerson(options)
	if testP.UPPs == "000000" {
		t.Error("MakePerson did not roll a UPP")
	}
}
