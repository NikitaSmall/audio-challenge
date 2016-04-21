package util

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("1234567890")

// RandStringRunes generates random string of numbers with provided length. Used for uuid purposes.
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}