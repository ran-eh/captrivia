// pkg/utils/random.go
package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ShuffleStrings shuffles a slice of strings.
func ShuffleStrings(slice []string) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// RandomInt returns a random integer between min and max.
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}