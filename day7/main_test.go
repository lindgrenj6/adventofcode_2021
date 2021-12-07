package main

import (
	"testing"
)

var (
	part1Answer = 342534
	part2Answer = 94004208
)

func TestPart1(t *testing.T) {
	out := Part1()
	if out != part1Answer {
		t.Errorf(`Wrong output for part 1, expected "%v" got "%v"`, part1Answer, out)
	}
}

func TestPart2(t *testing.T) {
	out := Part2()
	if out != part2Answer {
		t.Errorf(`Wrong output for part 2, expected "%v" got "%v"`, part2Answer, out)
	}
}

func TestFuelCalc(t *testing.T) {
	if fuelFor(5-1) != 10 {
		t.Errorf(`Wrong calc for (5-1)`)
	}

	if fuelFor(16-5) != 66 {
		t.Errorf(`Wrong calc for (16-5)`)
	}

	if fuelFor(5-0) != 15 {
		t.Errorf(`Wrong calc for (5-0)`)
	}
}
