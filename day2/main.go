package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

// var input = util.ParseFileToStringSlice("./testinput.txt")

var input = util.ParseFileToStringSlice("./input.txt")

var parsed []inst

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

type inst struct {
	pos string
	amt int
}

func parse() {
	raw := funk.Map(input, func(in string) inst {
		x := strings.Split(in, " ")
		amt, err := strconv.Atoi(x[1])
		if err != nil {
			panic("bad int: " + x[1])
		}

		return inst{
			pos: x[0],
			amt: amt,
		}
	})

	parsed = raw.([]inst)
}

func Part1() int {
	if parsed == nil {
		parse()
	}

	fpos, depth := 0, 0

	for _, v := range parsed {
		switch v.pos {
		case "forward":
			fpos += v.amt
		case "down":
			depth += v.amt
		case "up":
			depth -= v.amt
		default:
			panic("shouldn't have done this: " + v.pos)
		}
	}

	return fpos * depth
}

func Part2() int {
	if parsed == nil {
		parse()
	}

	fpos, depth, aim := 0, 0, 0

	for _, v := range parsed {
		switch v.pos {
		case "forward":
			fpos += v.amt
			depth += aim * v.amt
		case "down":
			aim += v.amt
		case "up":
			aim -= v.amt
		default:
			panic("shouldn't have done this: " + v.pos)
		}
	}

	return fpos * depth
}
