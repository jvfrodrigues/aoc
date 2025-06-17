package day01

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

const Title = "Day 1: Historian Hysteria"

var lines = utils.ReadFile("./data/day01/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func PartOne() int {
	total := 0
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		numLeft, err := strconv.ParseInt(strings.TrimSpace(numbers[0]), 0, 32)
		if err != nil {
			panic(err)
		}
		numRight, err := strconv.ParseInt(strings.TrimSpace(numbers[1]), 0, 32)
		if err != nil {
			panic(err)
		}
		left = append(left, int(numLeft))
		right = append(right, int(numRight))
	}
	slices.SortFunc(left, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(right, func(a, b int) int {
		return a - b
	})
	for idx := range left {
		total += int(math.Abs(float64(left[idx] - right[idx])))
	}
	return total
}

func PartTwo() int {
	score := 0
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		numLeft, err := strconv.ParseInt(strings.TrimSpace(numbers[0]), 0, 32)
		if err != nil {
			panic(err)
		}
		numRight, err := strconv.ParseInt(strings.TrimSpace(numbers[1]), 0, 32)
		if err != nil {
			panic(err)
		}
		left = append(left, int(numLeft))
		right = append(right, int(numRight))
	}
	rightMapped := make(map[int]int)
	for _, num := range right {
		value, ok := rightMapped[num]
		if !ok {
			rightMapped[num] = 1
		} else {
			rightMapped[num] = value + 1
		}
	}
	for _, num := range left {
		value, ok := rightMapped[num]
		if ok {
            score += num * value
		}
	}
	return score
}
