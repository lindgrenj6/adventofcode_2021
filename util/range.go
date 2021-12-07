package util

import "github.com/thoas/go-funk"

func RangeOf(start, end int) []int {
	r := make([]int, AbsValue(end-start))
	for i := 0; i < len(r); i++ {
		r[i] = i
	}

	return r
}

func ReverseRangeOf(start, end int) []int {
	return funk.ReverseInt(RangeOf(start, end))
}
