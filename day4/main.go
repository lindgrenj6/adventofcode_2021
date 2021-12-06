package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = util.ParseFileToSliceOfStringSlicesDoubleDelimited("./testinput.txt")

var input = util.ParseFileToSliceOfStringSlicesDoubleDelimited("./input.txt")

func parse() ([]int, [][][]int) {
	bingoBoards := make([][][]int, len(input[1:]))

	for idx, raw := range input[1:] {
		lines := strings.Split(raw, "\n")

		rawRow := funk.Map(lines, func(s string) []int {
			intStings := strings.Split(strings.ReplaceAll(strings.TrimSpace(s), "  ", " "), " ")
			ints := funk.Map(intStings, func(is string) int {
				i, err := strconv.Atoi(is)
				if err != nil {
					panic(err)
				}
				return i
			})

			return ints.([]int)
		})

		bingoBoards[idx] = rawRow.([][]int)
	}

	rawInstructions := strings.Split(input[0], ",")
	instructions := make([]int, len(rawInstructions))

	for i := 0; i < len(rawInstructions); i++ {
		parsed, err := strconv.Atoi(rawInstructions[i])
		if err != nil {
			panic(err)
		}
		instructions[i] = parsed
	}

	return instructions, bingoBoards
}

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

func Part1() int {
	inst, boards := parse()

	for i := 0; i < len(inst); i++ {
		for _, board := range boards {
			if checkBoard(inst[0:i], board) {
				longVec := funk.Reduce(board, func(acc, row []int) []int {
					return append(acc, row...)
				}, make([]int, 0))

				sumOfMissed := funk.SumInt(funk.FilterInt(longVec.([]int), func(j int) bool {
					return !funk.Contains(inst[0:i], j)
				}))

				return sumOfMissed * inst[i-1]
			}
		}
	}

	return -1
}

func Part2() int {
	inst, boards := parse()
	wins := make(map[int]bool)

	for i := 0; i < len(inst); i++ {
		for j := 0; j < len(boards); j++ {
			if wins[j] {
				continue
			}

			if checkBoard(inst[0:i], boards[j]) {
				wins[j] = true

				if len(wins) == len(boards) {
					longVec := funk.Reduce(boards[j], func(acc, row []int) []int {
						return append(acc, row...)
					}, make([]int, 0))

					filtered := funk.FilterInt(longVec.([]int), func(j int) bool {
						return !funk.Contains(inst[0:i], j)
					})

					sumOfMissed := funk.SumInt(filtered)
					return sumOfMissed * inst[i-1]
				}
			}
		}
	}

	return -1
}

func checkBoard(called []int, board [][]int) bool {
	for i := 0; i < len(board); i++ {
		if checkRow(called, board[i]) {
			return true
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if checkColumn(called, i, board) {
			return true
		}
	}

	return false
}

func checkRow(called []int, row []int) bool {
	return len(funk.InnerJoinInt(called, row)) == len(row)
}

func checkColumn(called []int, col int, board [][]int) bool {
	vec := make([]int, len(board))
	for i, val := range board {
		vec[i] = val[col]
	}

	return len(funk.InnerJoinInt(called, vec)) == len(vec)
}
