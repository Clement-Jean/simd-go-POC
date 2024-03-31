package tests

import (
	"math"
	"math/rand"
)

func genUint8Arr(l int) []uint8 {
	a := make([]uint8, l)
	rand.Read(a)
	return a
}

func genInt8Arr(l int) []int8 {
	a := make([]int8, l)
	for i := 0; i < len(a); i++ {
		a[i] = int8(rand.Intn(math.MaxInt8-math.MinInt8) + math.MinInt8)
	}
	return a
}
