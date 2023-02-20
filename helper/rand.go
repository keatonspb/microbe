package helper

import "math/rand"

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandInt(n int) int {
	return rand.Intn(n)
}

func RandFloat(n float64) float64 {
	return float64(rand.Int63n(int64(n)))
}
