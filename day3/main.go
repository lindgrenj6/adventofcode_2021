package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = util.ParseFileToSliceOfStringSlices("./testinput.txt")

var input = util.ParseFileToSliceOfStringSlices("./input.txt")

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	b := strings.Builder{}

	for col := 0; col < len(input[0]); col++ {
		common := mostCommon(col)
		if common["0"] > common["1"] {
			b.WriteString("0")
			continue
		}
		b.WriteString("1")
	}

	gamma, err := strconv.ParseUint(b.String(), 2, 64)
	if err != nil {
		panic(err)
	}

	size := (64 - len(input[0]))

	// bit hacking is literally my favorite thing
	epsilon := ^gamma << size >> size

	return int(gamma * epsilon)
}

func Part2() int {
	for i := 0; i < len(input[0]); i++ {
		selected := getSelection(mostCommon(i))

		filtered := funk.Filter(input, func(row []string) bool {
			if selected == "all" {
				return row[i] == "1"
			}
			return row[i] == selected
		})

		input = filtered.([][]string)
		if len(input) == 1 {
			break
		}
	}

	o2rating, err := strconv.ParseUint(strings.Join(input[0], ""), 2, 64)
	if err != nil {
		panic(err)
	}

	input = util.ParseFileToSliceOfStringSlices("./input.txt")

	for i := 0; i < len(input[0]); i++ {
		selected := getSelection(leastCommon(i))

		filtered := funk.Filter(input, func(row []string) bool {
			if selected == "all" {
				return row[i] == "0"
			}
			return row[i] == selected
		})

		input = filtered.([][]string)
		if len(input) == 1 {
			break
		}
	}

	scrubRating, err := strconv.ParseUint(strings.Join(input[0], ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(scrubRating) * int(o2rating)
}

func mostCommon(col int) map[string]int {
	counts := make(map[string]int)
	for row := 0; row < len(input); row++ {
		counts[input[row][col]]++
	}

	return counts
}

func leastCommon(col int) map[string]int {
	mapping := mostCommon(col)

	return map[string]int{
		"0": mapping["1"],
		"1": mapping["0"],
	}
}

func getSelection(mappings map[string]int) string {
	if mappings["0"] == mappings["1"] {
		return "all"
	}
	if mappings["0"] > mappings["1"] {
		return "0"
	}

	return "1"
}
