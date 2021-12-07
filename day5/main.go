package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = util.ParseFileToStringSlice("./testinput.txt")

var input = util.ParseFileToStringSlice("./input.txt")

type Vent struct {
	StartX, StartY int
	EndX, EndY     int
}

var parsed []Vent

func parse() []Vent {
	if parsed != nil {
		return parsed
	}

	vents := make([]Vent, len(input))

	ventRe := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for i, line := range input {
		matches := ventRe.FindAllStringSubmatch(line, -1)
		v := Vent{
			StartX: util.ParseToInt(matches[0][1]),
			StartY: util.ParseToInt(matches[0][2]),
			EndX:   util.ParseToInt(matches[0][3]),
			EndY:   util.ParseToInt(matches[0][4]),
		}

		vents[i] = v
	}

	parsed = vents
	return parsed
}

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	ventLines := parse()
	grid := make(map[string]int)

	for _, l := range ventLines {
		// only doing straight lines for p1
		if l.StartX != l.EndX && l.StartY != l.EndY {
			continue
		}

		switch {
		case l.StartX <= l.EndX && l.StartY <= l.EndY:
			for i := l.StartX; i <= l.EndX; i++ {
				for j := l.StartY; j <= l.EndY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX >= l.EndX && l.StartY <= l.EndY:
			for i := l.EndX; i <= l.StartX; i++ {
				for j := l.StartY; j <= l.EndY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX <= l.EndX && l.StartY >= l.EndY:
			for i := l.StartX; i <= l.EndX; i++ {
				for j := l.EndY; j <= l.StartY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX >= l.EndX && l.StartY >= l.EndY:
			for i := l.EndX; i <= l.StartX; i++ {
				for j := l.EndY; j <= l.StartY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		default:
			panic("unhandled case: " + fmt.Sprintf("%#v", l))
		}

	}

	return funk.Reduce(funk.Values(grid), func(acc, i int) int {
		if i >= 2 {
			acc++
		}

		return acc
	}, 0).(int)
}

func Part2() int {
	ventLines := parse()
	grid := make(map[string]int)

	for _, l := range ventLines {
		switch {
		case l.StartX <= l.EndX && l.StartY <= l.EndY:
			if !SlopeCheck(l) {
				continue
			}
			for i := l.StartX; i <= l.EndX; i++ {
				for j := l.StartY; j <= l.EndY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX >= l.EndX && l.StartY <= l.EndY:
			if !SlopeCheck(l) {
				continue
			}
			for i := l.EndX; i <= l.StartX; i++ {
				for j := l.StartY; j <= l.EndY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX <= l.EndX && l.StartY >= l.EndY:
			if !SlopeCheck(l) {
				continue
			}
			for i := l.StartX; i <= l.EndX; i++ {
				for j := l.EndY; j <= l.StartY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		case l.StartX >= l.EndX && l.StartY >= l.EndY:
			if !SlopeCheck(l) {
				continue
			}
			for i := l.EndX; i <= l.StartX; i++ {
				for j := l.EndY; j <= l.StartY; j++ {
					grid[fmt.Sprintf("%d,%d", i, j)]++
				}
			}
		default:
			panic("unhandled case: " + fmt.Sprintf("%#v", l))
		}

	}

	return funk.Reduce(funk.Values(grid), func(acc, i int) int {
		if i >= 2 {
			acc++
		}

		return acc
	}, 0).(int)
}

func SlopeCheck(v Vent) bool {
	if v.StartX != v.EndX && v.StartY != v.EndY {
		return true
	}

	slope := (math.Abs(float64((v.EndY - v.StartY)))) / (math.Abs(float64(v.EndX - v.StartX)))

	return slope == 1
}
