package main

import (
	"fmt"
	"sort"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

var input = util.ParseFileToSliceOfIntSlices("./input.txt")

// var input = util.ParseFileToSliceOfIntSlices("./testinput.txt")

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	nums := make([]int, 0)
	for col := 0; col < len(input[0]); col++ {
		for row := 0; row < len(input); row++ {
			if checkLowPoint(row, col) {
				nums = append(nums, input[row][col])
			}
		}
	}

	risks := funk.Map(nums, func(i int) int { return i + 1 })
	return funk.SumInt(risks.([]int))
}

func checkLowPoint(x, y int) bool {
	if x == 0 {
		if y == 0 {
			return input[x][y] < input[x][y+1] && input[x][y] < input[x+1][y]
		}
		if y == len(input[0])-1 {
			return input[x][y] < input[x][y-1] && input[x][y] < input[x+1][y]
		}
		return input[x][y] < input[x][y+1] && input[x][y] < input[x+1][y] && input[x][y] < input[x][y-1]
	}

	if x == len(input)-1 {
		if y == 0 {
			return input[x][y] < input[x][y+1] && input[x][y] < input[x-1][y]
		}
		if y == len(input[0])-1 {
			return input[x][y] < input[x][y-1] && input[x][y] < input[x-1][y]
		}
		return input[x][y] < input[x][y+1] && input[x][y] < input[x-1][y] && input[x][y] < input[x][y-1]
	}

	if y == 0 {
		return input[x][y] < input[x-1][y] && input[x][y] < input[x][y+1] && input[x][y] < input[x+1][y]
	}

	if y == len(input[0])-1 {
		return input[x][y] < input[x-1][y] && input[x][y] < input[x][y-1] && input[x][y] < input[x+1][y]
	}

	return input[x][y] < input[x-1][y] && input[x][y] < input[x+1][y] && input[x][y] < input[x][y-1] && input[x][y] < input[x][y+1]
}

type Point struct {
	X, Y  int
	Count int
}

func Part2() int {
	lowPoints := make([]Point, 0)

	for row := 0; row < len(input[0]); row++ {
		for col := 0; col < len(input); col++ {
			if checkLowPoint(col, row) {
				lowPoints = append(lowPoints, Point{X: row, Y: col})
			}
		}
	}

	for i := 0; i < len(lowPoints); i++ {
		found := countUntilWall(lowPoints[i].X, lowPoints[i].Y)
		lowPoints[i].Count = len(found)
	}
	fmt.Printf("lowPoints: %v\n", lowPoints)

	counts := funk.Map(lowPoints, func(p Point) int {
		return p.Count
	}).([]int)
	sort.Ints(counts)

	return funk.Reduce(counts[len(counts)-3:], func(acc, i int) int {
		return acc * i
	}, 1).(int)
}

func countUntilWall(x, y int) []Point {
	fmt.Printf("x, y: %v, %v\n", x, y)
	points := make([]Point, 0)
	for i := x; i < len(input[0])-1 || input[y][i] == 9; i++ {
		points = append(points, Point{X: i, Y: y})
	}
	for i := x; i > 0 || input[y][i] == 9; i-- {
		points = append(points, Point{X: i, Y: y})
	}
	for i := y; i < len(input)-1 || input[i][x] == 9; i++ {
		points = append(points, Point{X: x, Y: i})
	}
	for i := y; i > 0 || input[i][x] == 9; i-- {
		points = append(points, Point{X: x, Y: i})
	}

	fmt.Printf("len(points): %v\n", len(points))
	return points
}
