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

func ParseFileToStringSlice(filename string) []string {
	input := ReadFile(filename)
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

func ReadFile(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(input))
}

func ParseFileToSliceOfIntSlices(filename string) [][]int {
	lines := ParseFileToStringSlice(filename)

	raw := funk.Map(lines, func(line string) []int {
		return funk.Map(strings.Split(line, ""), func(n string) int {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic("bad int parse: " + n)
			}

			return i
		}).([]int)
	})

	return raw.([][]int)
}
