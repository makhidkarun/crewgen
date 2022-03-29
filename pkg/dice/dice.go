// dice.go

package dice

import (
	crand "crypto/rand"
	mbig "math/big"
)

// OneD6 returns a random int between 1 and 6, inclusive.
func OneD6() int {
	return Random(1, 6)
}

// TwoD6 returns a random int between 2 and 12, inclusive
func TwoD6() int {
	return Random(1, 6) + Random(1, 6)
}

// Random takes min and max ints and returns an int in the
// range min to max, inclusive.
func Random(min int, max int) int {
	spread := max - min + 1
	spread64 := int64(spread)
	num, _ := crand.Int(crand.Reader, mbig.NewInt(spread64))
	roll := int(num.Int64()) + min
	return roll
}
