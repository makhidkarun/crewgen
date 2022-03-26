package datamine_test

import (
	"os"
	"strings"
	"testing"

	"github.com/makhidkarun/crewgen/pkg/datamine"
)

var options map[string]string

func TestMain(m *testing.M) {
	options = make(map[string]string)
	options["game"] = "2d6"
	exitVal := m.Run()
	os.Exit(exitVal)
}

func hasElement(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TestStringInArray(t *testing.T) {
	var genders []string = []string{"F", "M"}
	option1 := "F"
	option2 := "R"
	if !datamine.StringInArray(option1, genders) {
		t.Error(`Missing an F.`)
	}
	if datamine.StringInArray(option2, genders) {
		t.Error(`Never met R.`)
	}
}

func TestFirstStringFromArray(t *testing.T) {
	var arr []string = []string{"first", "second"}
	expected := "first"
	result := datamine.FirstStringInArray(arr)
	if result != expected {
		t.Error(`TestFirstStringFromArray failed`)
	}
}

func TestRandomStringFromArray(t *testing.T) {
	var genders []string = []string{"F", "M"}
	gender := datamine.RandomStringFromArray(genders)
	if !datamine.StringInArray(gender, genders) {
		t.Error(`Bad gender output.`)
	}
	var ranks []string = []string{"PVT", "SGT", "LT", "CPT", "MAJ"}
	rank := datamine.RandomStringFromArray(ranks)
	if !datamine.StringInArray(rank, ranks) {
		t.Error(`Bad rank output.`)
	}
}

func TestGetFemaleFirstName(t *testing.T) {
	options["gender"] = "F"
	options["datadir"] = "data"
	options["lastName"] = ""
	name := datamine.GetName(options)
	if len(name) < 5 {
		t.Error(`Name too short`)
	}
}

func TestGetMaleFirstName(t *testing.T) {
	options["gender"] = "M"
	options["datadir"] = "data"
	options["lastName"] = ""
	name := datamine.GetName(options)
	if len(name) < 5 {
		t.Error(`Name too short`)
	}
}

func TestLastName(t *testing.T) {
	options["gender"] = "F"
	options["datadir"] = "testdata"
	options["lastName"] = "Domici"
	name := datamine.GetName(options)
	if !strings.Contains(name, options["lastName"]) {
		t.Error("In datamine, lastName does not match")
	}
}

func TestArrayFromFile(t *testing.T) {
	datafile := "testdata/human_female_first.txt"
	items := datamine.ArrayFromFile(datafile)
	if len(items) == 0 {
		t.Error("TestArrayFromFile has no items")
	}
	for _, item := range items {
		if strings.HasPrefix(item, "#") {
			t.Errorf("TestArrayFromFile let a pounder get in: %s", item)
		}
		if len(item) < 1 {
			t.Errorf("TestArrayFromFile let a blank line in: %s", item)
		}
	}
}

func TestLineToList(t *testing.T) {
	line := "Mechanic,Engineer,Electrical"
	sep := ","
	list := datamine.LineToList(line, sep)
	if len(list) < 3 {
		t.Errorf("TestLineToList has %d items\n", len(list))
	}
	if list[0] != "Mechanic" {
		t.Errorf("TestLineToList; expected Mechanic, got %q.\n", list[0])
	}
}

func TestDataFromListLine(t *testing.T) {
	var careerList []string
	careerOne := "navy:engineer:Mechanic,Engineer,Electrical"
	careerTwo := "army:infantry:GunCbt(CbtR),Brawling,Recon"
	careerList = append(careerList, careerOne)
	careerList = append(careerList, careerTwo)
	expected := "navy"
	sep := ":"
	career := datamine.DataFromListLine(careerList, expected, sep, 0)
	if career != "navy" {
		t.Errorf("TestDataFromListLine failed: wanted %s, got %q.\n", expected, career)
	}
	job := datamine.DataFromListLine(careerList, expected, sep, 1)
	if job != "engineer" {
		t.Errorf("TestDataFromListLine failed: wanted engineer, got %q.\n", job)
	}
	skillList := datamine.DataFromListLine(careerList, expected, sep, 2)
	if skillList != "Mechanic,Engineer,Electrical" {
		t.Errorf("TestDataFromListLine failed: wanted skillList, got %q.\n", skillList)
	}
}

func TestHeadersFromList(t *testing.T) {
	datafile := "testdata/careers.txt"
	items := datamine.ArrayFromFile(datafile)
	headers := datamine.HeadersFromList(items, ":")
	if len(headers) == 0 {
		t.Error("TestHeadersFromList has no items")
	}
	for _, item := range headers {
		if strings.HasPrefix(item, "#") {
			t.Error("TestHeadersFromList let a comment in")
		}
	}
	expectedList := []string{"navy", "army", "merchant", "marines", "scout", "other", "mercenary"}
	if len(headers) < len(expectedList) {
		t.Error("TestHeadersFromList has the wrong list count")
	}
	for _, header := range headers {
		if !hasElement(expectedList, header) {
			t.Errorf("TestHeadersFromList mixed expectedList does not have %s", header)
		}
	}
}

func TestCareerList(t *testing.T) {
	datafile := "testdata/careers.txt"
	careers := datamine.CareerList(datafile)
	if len(careers) == 0 {
		t.Error("TestCareerList has no items")
	}
	for _, item := range careers {
		if strings.HasPrefix(item, "#") {
			t.Error("TestCareerList let a comment in")
		}
	}
	expectedList := []string{"navy", "army", "merchant", "marines", "scout", "other", "mercenary"}
	if len(careers) < len(expectedList) {
		t.Error("TestCareerList has the wrong list count")
	}
	for _, career := range careers {
		if !hasElement(expectedList, career) {
			t.Errorf("TestCareerList does not have %s", career)
		}
	}
}

func TestCareerSkills(t *testing.T) {
	datafile := "testdata/careers.txt"
	career := "navy"
	careerSkills := datamine.CareerSkills(datafile, career)
	if len(careerSkills) == 0 {
		t.Error("TestCareerSkills has no careers")
	}
	if careerSkills[0] != "Mechanic" {
		t.Error("TestCareerSkills has the wrong list")
	}
}

func TestJobList(t *testing.T) {
	datafile := "testdata/jobs.txt"
	jobs := datamine.JobList(datafile)
	if len(jobs) == 0 {
		t.Error("TestJobList has no items")
	}
	for _, item := range jobs {
		if strings.HasPrefix(item, "#") {
			t.Error("TestJobList let a comment in")
		}
	}
	expectedList := []string{"infantry", "marine", "commando", "scout", "spacer", "merchant", "other", "pilot",
		"medic", "gunner", "navigator", "steward", "engineer"}
	if len(jobs) < len(expectedList) {
		t.Error("TestJobList has the wrong list count")
	}
	for _, job := range jobs {
		if !hasElement(expectedList, job) {
			t.Errorf("TestJobList does not have %s", job)
		}
	}
}

func TestJobSkillList(t *testing.T) {
	datafile := "testdata/jobs.txt"
	job := "other"
	jobSkills := datamine.JobSkills(datafile, job)
	if len(jobSkills) == 0 {
		t.Error("TestJobSkillList has no jobs")
	}
	if jobSkills[0] != "Streetwise" {
		t.Error("TestJobSkillList has the wrong list")
	}
}

func TestDefaultJob(t *testing.T) {
	datafile := "testdata/careers.txt"
	result := datamine.DefaultJob(datafile, "scout")
	expected := "scout"
	if result != expected {
		t.Errorf("TestDefaultJob got %s, expected %s", result, expected)
	}
}

func TestOptions(t *testing.T) {
	careerfile := "testdata/careers.txt"
	jobfile := "testdata/jobs.txt"

	result := datamine.ListOptions(careerfile, jobfile)
	if !strings.Contains(result, "Scout") {
		t.Errorf("result does not contain Scout: %s", result)
	}

}
