package random

import (
	"math"
	"math/rand"
	"time"
)

func Seed(seed int64) {
	rand.Seed(seed)
}

func RandSeed() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func Float() float64 {
	return rand.Float64()
}

func Int() int {
	return rand.Int()
}

func FloatRange(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func IntRange(min int, max int) int {
	randRange := float64(max - min)
	return min + int(math.Round(rand.Float64()*randRange))
}

func Boolean() bool {
	return WeightedBool(0.5)
}

func WeightedBool(weight float64) bool {
	return rand.Float64() < weight
}
