package main

import (
	"fmt"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = funk.Map(strings.Split(util.ReadFile("./testinput.txt"), ","), func(s string) int {
// 	return util.ParseToInt(s)
// }).([]int)

var input = funk.Map(strings.Split(util.ReadFile("./input.txt"), ","), func(s string) int {
	return util.ParseToInt(s)
}).([]int)

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	feesh := make([]int, len(input))
	copy(feesh, input)
	for i := 0; i < 80; i++ {
		newFeesh := make([]int, 0, 0)

		for j := 0; j < len(feesh); j++ {
			// spawn feesh
			if feesh[j] == 0 {
				feesh[j] = 6
				newFeesh = append(newFeesh, 8)
			} else {
				feesh[j]--
			}
		}

		feesh = append(feesh, newFeesh...)
	}

	return len(feesh)
}

func Part2() int {
	feesh := make([]int, 9)
	funk.ForEach(input, func(f int) {
		feesh[f]++
	})

	for day := 0; day < 256; day++ {
		newFeesh := 0
		newFeesh, feesh = feesh[0], feesh[1:]
		feesh = append(feesh, newFeesh)
		feesh[6] += newFeesh
	}

	return funk.SumInt(feesh)
}
