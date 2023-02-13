package helpers

import (
	"math/rand"
	"time"
)

// IsTrue returns true/false randomly
func IsTrue() bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	res := r1.Intn(2)
	return res == 1
}
