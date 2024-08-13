package goevm

import (
	"math"
)

func CalcMemoryGasCost(size uint64) uint64 {
	word := (size + 31) / 32
	cost := (uint64(math.Pow(float64(word), 2)))/512 + (3 * word)
	return cost
}
