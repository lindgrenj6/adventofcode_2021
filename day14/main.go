package main

import (
	"fmt"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = util.ParseFileToStringSlice("./testinput.txt")

var input = util.ParseFileToStringSlice("./input.txt")

type Parsed struct {
	Template []string
	Mappings map[[2]string]string // using an array since they are comparable
}

var parsed *Parsed

func parse() *Parsed {
	if parsed != nil {
		return parsed
	}

	p := Parsed{Template: strings.Split(input[0], ""), Mappings: make(map[[2]string]string)}

	funk.ForEach(input[2:], func(s string) {
		var key [2]string
		parts := strings.Fields(s)
		// hacky stuff to convince copy that the array is a slice
		copy(key[:], strings.Split(parts[0], "")[0:2])

		p.Mappings[key] = parts[2]
	})

	parsed = &p
	return &p
}

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	parsed := parse()
	polymer := parsed.Template

	for count := 0; count < 10; count++ {
		newLine := make([]string, 0)

		for i := 0; i < len(polymer)-1; i++ {
			newLine = append(newLine, polymer[i])
			newLine = append(newLine, parsed.Mappings[[2]string{polymer[i], polymer[i+1]}])
		}

		// tack on the last element - otherwise it gets forgotten.
		polymer = append(newLine, polymer[len(polymer)-1])
	}

	counts := funk.Reduce(polymer, func(acc map[string]int, p string) map[string]int {
		acc[p]++
		return acc
	}, make(map[string]int)).(map[string]int)

	vals := funk.Values(counts).([]int)
	return funk.MaxInt(vals) - funk.MinInt(vals)
}

func Part2() int {
	parsed := parse()
	polymerMap := make(map[[2]string]int)

	for i := 0; i < len(parsed.Template)-1; i++ {
		k := [2]string{parsed.Template[i], parsed.Template[i+1]}
		polymerMap[[2]string{parsed.Template[i], parsed.Mappings[k]}]++
	}
	fmt.Printf("polymerMap: %v\n", polymerMap)

	for i := 0; i < 40; i++ {
	}
	return -1
}
