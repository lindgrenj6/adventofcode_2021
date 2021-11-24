package main

import (
	"fmt"

	"github.com/lindgrenj6/adventofcode_2021/util"
)

var input = util.ParseFileToIntSlice("./input.txt")

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	return input[0]
}

func Part2() int {
	return input[1]
}
