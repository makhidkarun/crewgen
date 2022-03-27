package person_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/makhidkarun/crewgen/pkg/person"
)

func hasElement(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

var options map[string]string

func TestMain(m *testing.M) {
	options = make(map[string]string)
	options["terms"] = "1"
	options["career"] = "navy"
	options["datadir"] = "testdata"
	options["careerFile"] = path.Join(options["datadir"], "careers.txt")
	options["jobFile"] = path.Join(options["datadir"], "jobs.txt")
	options["game"] = "2d6"
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestMakePerson(t *testing.T) {
	testP := person.MakePerson(options)
	var tP interface{} = testP.Name
	if data, ok := tP.(string); !ok {
		t.Errorf("MakePerson failed by name test %s", data)
	}
}

func TestMakePersonAge(t *testing.T) {
	options["terms"] = "1"
	options["game"] = "2d6"
	testP := person.MakePerson(options)
	if testP.Age <= 21 || testP.Age >= 26 {
		t.Errorf("MakePerson failed age for 2d6: %d", testP.Age)
	}
}

func TestMakePersonAge10(t *testing.T) {
	options["terms"] = "10"
	options["game"] = "2d6"
	testP := person.MakePerson(options)
	if testP.Age < 58 || testP.Age > 61 {
		t.Errorf("MakePerson failed age for 10 terms 2d6: %d", testP.Age)
	}
}

func TestMakePersonAgeBRP(t *testing.T) {
	options["terms"] = "10"
	options["game"] = "brp"
	testP := person.MakePerson(options)
	if testP.Age < 18 || testP.Age > 21 {
		t.Error("MakePerson failed age for BRP")
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
	options["career"] = "navy"
	testP := person.MakePerson(options)
	if testP.Career != "navy" {
		t.Error(`MakePerson failed to specify Navy career`)
	}
}

func TestCareerMerchant(t *testing.T) {
	options["career"] = "merchant"
	testP := person.MakePerson(options)
	if testP.Career != "merchant" {
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

func TestUPPsEmptyUPP(t *testing.T) {
	testP := person.MakePerson(options)
	if testP.UPPs == "000000" {
		t.Error("MakePerson did not roll a UPP")
	}
}

func Test2d6Stats(t *testing.T) {
	options["game"] = "2d6"
	testP := person.MakePerson(options)
	for _, value := range testP.UPP {
		if value < 2 || value > 15 {
			t.Errorf("MakePerson has a UPP value outside of 2-15: %d", value)
		}
	}
}

func TestBRPStats(t *testing.T) {
	options["game"] = "brp"
	testP := person.MakePerson(options)
	for _, value := range testP.UPP {
		if value < 3 || value > 18 {
			t.Errorf("MakePerson has a BRP UPP value outside of 3-18: %d", value)
		}
	}
}

func TestSkills(t *testing.T) {
	options["terms"] = "4"
	options["job"] = "other"
	testP := person.MakePerson(options)
	if len(testP.SkillString) < 8 {
		t.Error(`MakePerson failed to specify a long skillstring`)
	}
}

func TestSkillsStartMinus(t *testing.T) {
	testP := person.MakePerson(options)
	if strings.HasPrefix(testP.SkillString, "-") {
		t.Errorf("SkillString starting with a '-': %s\n", testP.SkillString)
	}
}

func TestNoJob(t *testing.T) {
	options["terms"] = "4"
	options["job"] = ""
	testP := person.MakePerson(options)
	if len(testP.SkillString) <= 0 {
		t.Errorf("MakePerson did a blank skillstring for no job: %q\n", testP.SkillString)
	}
}

func TestDefaultJob(t *testing.T) {
	options["career"] = "scout"
	testP := person.MakePerson(options)
	if !strings.Contains(testP.SkillString, "Pilot-") {
		t.Errorf("TestDefaultJob does not give Scouts pilot: %s", testP.SkillString)
	}
}

func TestMercenaryCareer(t *testing.T) {
	options["career"] = "mercenary"
	testP := person.MakePerson(options)
	if len(testP.SkillString) < 8 {
		t.Error(`TestMercenaryCareer failed to specify a long skillstring`)
	}
	if testP.Career != "mercenary" {
		t.Error(`TestMercenaryCareer failed to specify Mercenary career`)
	}
	if !strings.Contains(testP.SkillString, "GunCbt(CbtR)") {
		t.Errorf("TestMercenaryCareer does not give infantry CbtR: %q in %q", testP.SkillString, testP.Career)
	}
}

func TestLastName(t *testing.T) {
	options["lastName"] = "Domici"
	testP := person.MakePerson(options)
	if !strings.Contains(testP.Name, "Domici") {
		t.Error("In person, TestLastName does not match")
	}
}

func TestSetGenderF(t *testing.T) {
	options["gender"] = "F"
	testP := person.MakePerson(options)
	if !strings.Contains(testP.Gender, "F") {
		t.Errorf("In person, TestSetGender failed for F, got %s", testP.Gender)
	}
}

func TestSetGenderM(t *testing.T) {
	options["gender"] = "M"
	testP := person.MakePerson(options)
	if !strings.Contains(testP.Gender, "M") {
		t.Errorf("In person, TestSetGender failed for M, got %s", testP.Gender)
	}
}

func TestGetPlot(t *testing.T) {
	//plotfile := path.Join(options["datadir"], "plots.txt")
	plots := []string{"Love", "Money", "Murder"}
	testP := person.MakePerson(options)
	if !hasElement(plots, testP.Plot) {
		t.Errorf("In person, TestGetPlot gave %s, not an option\n", testP.Plot)
	}
}
