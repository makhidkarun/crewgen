package tools_test

import (
	mrand "math/rand"

	"github.com/makhidkarun/crewgen/pkg/tools"
	"testing"
)

var rng = mrand.New(mrand.NewSource(99))

func TestOneD6(t *testing.T) {
	roll := tools.OneD6()
	if roll < 1 || roll > 6 {
		t.Error(`OneD6 failed`)
	}
}

func TestTwoD6(t *testing.T) {
	roll := tools.TwoD6()
	if roll < 2 || roll > 12 {
		t.Error(`TwoD6 failed`)
	}
}

func TestAge(t *testing.T) {
	age := tools.Age(2)
	if age < 26 || age > 29 {
		t.Error(`Age failed`)
	}
}

func TestFormatUPP(t *testing.T) {
	upp := [6]int{7, 7, 15, 9, 12, 12}
	newUPP := tools.FormatUPP(upp)
	if newUPP != "77F9CC" {
		t.Error(`FormatUPP failed.`)
	}
}

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
