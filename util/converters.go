package util

import "strconv"

func ParseToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("bad string: " + s)
	}

	return i
}
