package main

import (
	"fmt"
	"sort"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

var input = util.ParseFileToSliceOfStringSlices("./input.txt")

// var input = util.ParseFileToSliceOfStringSlices("./testinput.txt")

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

var openers = []string{"[", "(", "{", "<"}

func Part1() int {
	errors := make([]string, 0)

	for _, line := range input {
		st := util.NewStack()
		for _, s := range line {
			if funk.ContainsString(openers, s) {
				st.Push(s)
				continue
			}

			found := false
			switch s {
			case "]":
				if st.Peek() != "[" {
					errors = append(errors, s)
					found = true
				}
			case ")":
				if st.Peek() != "(" {
					errors = append(errors, s)
					found = true
				}
			case "}":
				if st.Peek() != "{" {
					errors = append(errors, s)
					found = true
				}
			case ">":
				if st.Peek() != "<" {
					errors = append(errors, s)
					found = true
				}
			}

			if found {
				break
			}

			st.Pop()
		}
	}

	return funk.Reduce(funk.Compact(errors), func(acc int, e string) int {
		switch e {
		case ")":
			acc += 3
		case "]":
			acc += 57
		case "}":
			acc += 1197
		case ">":
			acc += 25137
		default:
			panic("bad thing: " + e)
		}

		return acc
	}, 0).(int)
}

func Part2() int {
	corrupted := funk.Filter(input, isCorrupted).([][]string)
	additions := make([][]string, 0)

	for _, line := range corrupted {
		st := util.NewStack()
		for _, s := range line {
			if funk.ContainsString(openers, s) {
				st.Push(s)
				continue
			}

			switch s {
			case "]":
				if st.Peek() == "[" {
					st.Pop()
				}
			case ")":
				if st.Peek() == "(" {
					st.Pop()
				}
			case "}":
				if st.Peek() == "{" {
					st.Pop()
				}
			case ">":
				if st.Peek() == "<" {
					st.Pop()
				}
			}
		}

		completion := make([]string, 0)

		for {
			val := st.Pop()
			switch val {
			case "(":
				completion = append(completion, ")")
			case "{":
				completion = append(completion, "}")
			case "[":
				completion = append(completion, "]")
			case "<":
				completion = append(completion, ">")
			}

			if val == "" {
				break
			}
		}

		additions = append(additions, completion)
	}

	scores := funk.Map(additions, func(e []string) int {
		return funk.Reduce(e, func(acc int, s string) int {
			switch s {
			case ")":
				return acc*5 + 1
			case "]":
				return acc*5 + 2
			case "}":
				return acc*5 + 3
			case ">":
				return acc*5 + 4
			default:
				panic("bad thing: " + s)
			}
		}, 0).(int)
	}).([]int)

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func isCorrupted(line []string) bool {
	st := util.NewStack()
	for _, s := range line {
		if funk.ContainsString(openers, s) {
			st.Push(s)
			continue
		}

		switch s {
		case "]":
			if st.Peek() != "[" {
				return false
			}
		case ")":
			if st.Peek() != "(" {
				return false
			}
		case "}":
			if st.Peek() != "{" {
				return false
			}
		case ">":
			if st.Peek() != "<" {
				return false
			}
		}
		st.Pop()
	}

	return true
}
