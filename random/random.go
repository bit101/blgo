package random

import (
	"math"
	"math/rand"
	"time"
)

// Seed sets the prng seed.
func Seed(seed int64) {
	rand.Seed(seed)
}

// RandSeed seeds the prng with a random seed.
func RandSeed() int64 {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)
	return seed
}

// Float returns a random float from 0.0 to 1.0.
func Float() float64 {
	return rand.Float64()
}

// Int returns a random integer.
func Int() int {
	return rand.Int()
}

// FloatRange returns a random float from min to max.
func FloatRange(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// IntRange returns a random int from min to max.
func IntRange(min int, max int) int {
	randRange := float64(max - min)
	return min + int(math.Round(rand.Float64()*randRange))
}

// Boolean returns a random boolean.
func Boolean() bool {
	return WeightedBool(0.5)
}

// WeightedBool returns a weighted boolean.
func WeightedBool(weight float64) bool {
	return rand.Float64() < weight
}
