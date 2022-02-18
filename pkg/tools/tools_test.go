package tools_test

import (
	"github.com/makhidkarun/crewgen/pkg/tools"
	"testing"
)

func TestStringInArray(t *testing.T) {
	var genders []string = []string{"F", "M"}
	option1 := "F"
	option2 := "R"
	if !tools.StringInArray(option1, genders) {
		t.Error(`Missing an F.`)
	}
	if tools.StringInArray(option2, genders) {
		t.Error(`Never met R.`)
	}
}

func TestRandomStringFromArray(t *testing.T) {
	var genders []string = []string{"F", "M"}
	gender := tools.RandomStringFromArray(genders)
	if !tools.StringInArray(gender, genders) {
		t.Error(`Bad gender output.`)
	}
	var ranks []string = []string{"PVT", "SGT", "LT", "CPT", "MAJ"}
	rank := tools.RandomStringFromArray(ranks)
	if !tools.StringInArray(rank, ranks) {
		t.Error(`Bad rank output.`)
	}
}
