package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

var input = funk.Map(strings.Split(util.ReadFile("./input.txt"), ","), func(s string) int { return util.ParseToInt(s) }).([]int)

// var input = funk.Map(strings.Split(util.ReadFile("./testinput.txt"), ","), func(s string) int { return util.ParseToInt(s) }).([]int)

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	min, max := funk.MinInt(input), funk.MaxInt(input)
	lowest := math.MaxInt32

	for try := min; try <= max; try++ {
		fuel := 0
		for i := 0; i < len(input); i++ {
			fuel += util.AbsValue(input[i] - try)
		}

		if lowest > fuel {
			lowest = fuel
		}
	}

	return lowest
}

func Part2() int {
	min, max := funk.MinInt(input), funk.MaxInt(input)
	lowest := math.MaxInt32

	for try := min; try <= max; try++ {
		fuel := 0
		for i := 0; i < len(input); i++ {
			fuel += fuelFor(util.AbsValue(input[i] - try))
		}

		if lowest > fuel {
			lowest = fuel
		}
	}

	return lowest
}

func fuelFor(steps int) int {
	count := 0
	for i := 1; i <= steps; i++ {
		count += i
	}

	return count
}
