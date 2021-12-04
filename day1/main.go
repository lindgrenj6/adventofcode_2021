package main

import (
	"fmt"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

var input = util.ParseFileToIntSlice("./input.txt")

// var input = util.ParseFileToIntSlice("./testinput.txt")

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	count := 0
	last := input[0]

	for i := 0; i < len(input); i++ {
		if input[i] > last {
			count++
		}

		last = input[i]
	}

	return count
}

func Part2() int {
	count := 0
	// cheating with funk a bit here O-O
	last := funk.SumInt(input[0:3])

	for i := 1; i+3 <= len(input); i++ {
		if funk.SumInt(input[i:i+3]) > last {
			count++
		}

		last = funk.SumInt(input[i : i+3])
	}

	return count
}
