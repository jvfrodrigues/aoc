package day03

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const Title = "Day 3: Mull It Over"

var lines = utils.ReadFile("./data/day03/input.txt")

var pattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func PartOne() int {
	result := 0
	for _, line := range lines {
		matched := pattern.FindAll([]byte(line), -1)
		for _, match := range matched {
			strMatch := string(match)
			numbers := strings.Split(strings.Trim(strMatch, "mul()"), ",")
			mult := 1
			for _, number := range numbers {
				parsedNum, err := strconv.ParseInt(number, 0, 32)
				if err != nil {
					panic(err)
				}
				mult = mult * int(parsedNum)
			}
			result += mult
		}
	}
	return result
}

var doPattern = regexp.MustCompile(`don?'?t?\(\)`)

func PartTwo() int {
	result := 0
	active := true
	for _, line := range lines {
		doCommands := doPattern.FindAllIndex([]byte(line), -1)
		matched := pattern.FindAllIndex([]byte(line), -1)
		allMatches := append(doCommands, matched...)
		slices.SortFunc(allMatches, func(a, b []int) int {
			return a[0] - b[0]
		})
		for _, match := range allMatches {
			strMatch := line[match[0]:match[1]]
			if strMatch == "do()" {
				active = true
				continue
			} else if strMatch == "don't()" {
				active = false
				continue
			} else {
				if active {
					numbers := strings.Split(strings.Trim(strMatch, "mul()"), ",")
					mult := 1
					for _, number := range numbers {
						parsedNum, err := strconv.ParseInt(number, 0, 32)
						if err != nil {
							panic(err)
						}
						mult = mult * int(parsedNum)
					}
					result += mult
				}
			}
		}
	}
	return result
}
