package main

import (
	"testing"
)

var (
	part1Answer = 537
	part2Answer = 2
)

func TestPart1(t *testing.T) {
	out := Part1()
	if out != part1Answer {
		t.Errorf(`Wrong output for part 1, expected "%v" got "%v"`, part1Answer, out)
	}
}

func TestPart2(t *testing.T) {
	t.Skip("not completed")
	out := Part2()
	if out != part2Answer {
		t.Errorf(`Wrong output for part 2, expected "%v" got "%v"`, part2Answer, out)
	}
}
