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
