package main

import (
	"fmt"
	"strings"

	"github.com/lindgrenj6/adventofcode_2021/util"
	"github.com/thoas/go-funk"
)

var input = util.ParseFileToStringSlice("./input.txt")

// var input = util.ParseFileToStringSlice("./testinput.txt")
var parsed []entry

func main() {
	fmt.Printf("Part 1: %#v\n", Part1())
	fmt.Printf("Part 2: %#v\n", Part2())
}

type entry struct {
	signal []string
	output []string
}

func parse() []entry {
	if parsed != nil {
		return parsed
	}

	signals := make([]entry, len(input))

	for i, line := range input {
		parts := strings.Split(line, "|")
		signals[i] = entry{
			signal: strings.Fields(strings.TrimSpace(parts[0])),
			output: strings.Fields(strings.TrimSpace(parts[1])),
		}
	}

	parsed = signals
	return signals
}

func Part1() int {
	signals := parse()
	count := 0

	for _, signal := range signals {
		good := funk.FilterString(signal.output, func(s string) bool {
			leds := strings.Split(s, "")

			switch len(leds) {
			case 4:
				return true
			case 3:
				return true
			case 7:
				return true
			case 2:
				return true
			default:
				return false
			}
		})

		count += len(good)
	}

	return count
}

func Part2() int {
	signals := parse()
	perms := permutations("abcdefg")
	count := 0

	for _, s := range signals {
		nums := funk.Map(s.output, func(d string) string {
			chars := strings.Split(d, "")

			for i := 0; i < len(perms); i++ {
				mapping := generateMapping(perms[i])
				leds := funk.Map(chars, func(c string) string { return mapping[c] })

				d, found := findDigit(leds.([]string))
				if found {
					return d
				}
			}

			panic("should never happen, brute force always wins...eventually")
		})
		count += util.ParseToInt(strings.Join(nums.([]string), ""))
	}

	return count
}

var digits = map[string][]string{
	"0": {"top", "upleft", "upright", "lowerleft", "lowerright", "bottom"},
	"1": {"upright", "lowerright"},
	"2": {"top", "upright", "center", "lowerleft", "bottom"},
	"3": {"top", "upright", "center", "lowerright", "bottom"},
	"4": {"upleft", "upright", "center", "lowerright"},
	"5": {"top", "upleft", "center", "lowerright", "bottom"},
	"6": {"top", "upleft", "center", "lowerleft", "lowerright", "bottom"},
	"7": {"top", "upright", "lowerright"},
	"8": {"top", "upleft", "upright", "center", "lowerleft", "lowerright", "bottom"},
	"9": {"top", "upleft", "upright", "center", "lowerright", "bottom"},
}

func generateMapping(try string) map[string]string {
	return map[string]string{
		string(try[0]): "top",
		string(try[1]): "upleft",
		string(try[2]): "upright",
		string(try[3]): "center",
		string(try[4]): "lowerleft",
		string(try[5]): "lowerright",
		string(try[6]): "bottom",
	}
}

func findDigit(leds []string) (string, bool) {
	for digit, mapping := range digits {
		match := funk.IntersectString(leds, mapping)
		if len(match) == len(mapping) && len(match) == len(leds) {
			return digit, true
		}
	}

	return "", false
}

// stolen:
// https://www.golangprograms.com/golang-program-to-print-all-permutations-of-a-given-string.html
func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}
