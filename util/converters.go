package util

import (
	"math"
	"strconv"
)

func ParseToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("bad string: " + s)
	}

	return i
}

func AbsValue(i int) int {
	return int(math.Abs(float64(i)))
}
