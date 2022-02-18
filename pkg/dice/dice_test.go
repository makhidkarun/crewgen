// dice_test.go

package dice_test

import (
	//  mrand "math/rand"

	"github.com/makhidkarun/crewgen/pkg/dice"
	"testing"
)

//var rng = mrand.New(mrand.NewSouce(99))

func TestOneD6(t *testing.T) {
	roll := dice.OneD6()
	if roll < 1 || roll > 6 {
		t.Error(`OneD6 failed`)
	}
}

func TestTwoD6(t *testing.T) {
	roll := dice.TwoD6()
	if roll < 2 || roll > 12 {
		t.Error(`TwoD6 failed`)
	}
}
