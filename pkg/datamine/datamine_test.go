package datamine_test

import (
	"strings"

	"github.com/makhidkarun/crewgen/pkg/datamine"
	"testing"
)

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
	gender := "F"
	datadir := "data"
	name := datamine.GetName(gender, datadir)
	if len(name) < 5 {
		t.Error(`Name too short`)
	}
}

func TestGetMaleFirstName(t *testing.T) {
	gender := "M"
	datadir := "data"
	name := datamine.GetName(gender, datadir)
	if len(name) < 5 {
		t.Error(`Name too short`)
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
	careerOne := "Navy:engineer:Mechanic,Engineer,Electrical"
	careerTwo := "Army:infantry:GunCbt(CbtR),Brawling,Recon"
	careerList = append(careerList, careerOne)
	careerList = append(careerList, careerTwo)
	expected := "Navy"
	sep := ":"
	career := datamine.DataFromListLine(careerList, expected, sep, 0)
	if career != "Navy" {
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
	expectedList := []string{"Navy", "Army", "Merchant", "Marines", "Scout", "Other", "Mercenary"}
	if len(headers) != len(expectedList) {
		t.Error("TestHeadersFromList has the wrong list count")
	}
	for index, header := range headers {
		if expectedList[index] != header {
			t.Errorf("TestHeadersFromList mixed up %s and %s", expectedList[index], header)
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
	expectedList := []string{"Navy", "Army", "Merchant", "Marines", "Scout", "Other", "Mercenary"}
	if len(careers) != len(expectedList) {
		t.Error("TestCareerList has the wrong list count")
	}
	for index, career := range careers {
		if expectedList[index] != career {
			t.Errorf("TestCareerList mixed up %s and %s", expectedList[index], career)
		}
	}
}

func TestCareerSkills(t *testing.T) {
	datafile := "testdata/careers.txt"
	career := "Navy"
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
	if len(jobs) != len(expectedList) {
		t.Error("TestJobList has the wrong list count")
	}
	for index, job := range jobs {
		if expectedList[index] != job {
			t.Errorf("TestJobList mixed up %s and %s", expectedList[index], job)
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
	result := datamine.DefaultJob(datafile, "Scout")
	expected := "scout"
	if result != expected {
		t.Errorf("TestDefaultJob got %s, expected %s", result, expected)
	}
}
