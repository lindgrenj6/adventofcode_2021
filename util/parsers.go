package util

import (
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func ParseFileToIntSlice(filename string) []int {
	lines := ParseFileToStringSlice(filename)
	out := funk.Map(lines, func(str string) int {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		return i
	})

	return out.([]int)
}

func DoubleParseIntSlice(filename string) [][]string {
	input := readFile(filename)
	lines := strings.Split(strings.TrimSpace(input), "\n\n")

	return splitLines(lines)
}

func ParseFileToStringSlice(filename string) []string {
	input := readFile(filename)
	return strings.Split(strings.TrimSpace(input), "\n")
}

func ParseFileToSliceOfStringSlices(filename string) [][]string {
	lines := ParseFileToStringSlice(filename)

	return splitLines(lines)
}

func splitLines(lines []string) [][]string {
	out := funk.Map(lines, func(str string) []string {
		return strings.Split(str, "")
	})

	return out.([][]string)
}

func readFile(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(input)
}
