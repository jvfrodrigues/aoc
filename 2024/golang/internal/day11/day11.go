package day11

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"math"
	"strconv"
	"strings"
)

const Title = "Day 11: Plutonian Pebbles"

var lines = utils.ReadFile("./data/day11/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func PartOne() int {
	stones := readLines(lines)
	result := 0
	iteratedStones := iterateStones(stones, 25)
	for _, count := range iteratedStones {
		result += count
	}
	return result
}

func PartTwo() int {
	stones := readLines(lines)
	result := 0
	iteratedStones := iterateStones(stones, 75)
	for _, count := range iteratedStones {
		result += count
	}
	return result
}

func iterateStones(stones map[int64]int, iterations int) map[int64]int {
	endStones := make(map[int64]int)
	for key, value := range stones {
		endStones[key] = value
	}
	for i := 0; i < iterations; i++ {
		updatedStones := make(map[int64]int)
		for stone, count := range endStones {
			stringValue := strconv.FormatInt(stone, 10)
			valueLen := len(stringValue)
			if stone == 0 {
				_, ok := updatedStones[1]
				if !ok {
					updatedStones[1] = 0
				}
				updatedStones[1] += count
			} else if valueLen%2 == 0 {
				divisor := int64(math.Pow(10, float64(valueLen/2)))
				left := stone / divisor
				right := stone % divisor
				_, ok := updatedStones[left]
				if !ok {
					updatedStones[left] = 0
				}
				_, ok = updatedStones[right]
				if !ok {
					updatedStones[right] = 0
				}
				updatedStones[left] += count
				updatedStones[right] += count
			} else {
				newValue := stone * 2024
				_, ok := updatedStones[newValue]
				if !ok {
					updatedStones[newValue] = 0
				}
				updatedStones[newValue] += count
			}
		}
		endStones = updatedStones
	}
	return endStones
}

func readLines(lines []string) map[int64]int {
	stones := make(map[int64]int)
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			number = strings.TrimSpace(number)
			parsedNumber, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				panic(err)
			}
			_, ok := stones[parsedNumber]
			if !ok {
				stones[parsedNumber] = 0
			}
			stones[parsedNumber]++
		}
	}
	return stones
}
