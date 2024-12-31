package day07

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

const Title = "Day 7: Bridge Repair"

var lines = utils.ReadFile("./data/day07/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int64]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

var signsP1 = map[string]string{
	"0": "+",
	"1": "*",
}

func PartOne() int64 {
	groups, totals := readLines(lines)
	var total int64 = 0
	c := make(chan int64, len(groups))
	wg := sync.WaitGroup{}
	for idx, group := range groups {
		currTotal := totals[idx]
        wg.Add(1)
		go func() {
			defer wg.Done()
			isCalibrated := evaluateP1(group, currTotal)
			if isCalibrated {
				c <- currTotal
			}
		}()
	}
	wg.Wait()
	close(c)
	for foundTotal := range c {
		total += foundTotal
	}
	return total
}

var signsP2 = map[string]string{
	"0": "+",
	"1": "*",
	"2": "||",
}

func PartTwo() int64 {
	groups, totals := readLines(lines)
	c := make(chan int64, len(groups))
	wg := sync.WaitGroup{}
	var total int64 = 0
	for idx, group := range groups {
		currTotal := totals[idx]
		wg.Add(1)
		go func() {
			defer wg.Done()
			isCalibrated := evaluateP2(group, currTotal, group[0], 1)
			if isCalibrated {
				c <- currTotal
			}
		}()
	}
	wg.Wait()
	close(c)
	for foundTotal := range c {
		total += foundTotal
	}
	return total
}

func evaluateP1(group []int64, total int64) bool {
	possibilities := int(math.Pow(2, float64((len(group) - 1))))
	for i := 0; i < possibilities; i++ {
		binary := fmt.Sprintf("%0*b", len(group)-1, i)
		result := group[0]
		for idx, element := range binary {
			currOperation := signsP1[string(element)]
			if currOperation == "*" {
				result = result * group[idx+1]
			} else {
				result = result + group[idx+1]
			}
			if result > total {
				break
			}
		}
		if result == total {
			return true
		}
	}
	return false
}

func evaluateP2(group []int64, total int64, result int64, depth int) bool {
	if depth >= len(group) {
		return false
	}
	curr := group[depth]
	for _, sign := range signsP2 {
		if sign == "*" {
			currResult := result * curr
			if currResult == total && depth == len(group)-1 {
				return true
			} else {
				isValid := evaluateP2(group, total, currResult, depth+1)
				if isValid {
					return true
				}
			}
		} else if sign == "+" {
			currResult := result + curr
			if currResult == total && depth == len(group)-1 {
				return true
			} else {
				isValid := evaluateP2(group, total, currResult, depth+1)
				if isValid {
					return true
				}
			}
		} else {
			combinedValue := strconv.FormatInt(result, 10) + strconv.FormatInt(curr, 10)
			currResult, err := strconv.ParseInt(combinedValue, 0, 64)
			if err != nil {
				return false
			}
			if currResult == total && depth == len(group)-1 {
				return true
			} else {
				isValid := evaluateP2(group, total, currResult, depth+1)
				if isValid {
					return true
				}
			}
		}
	}
	return false
}

func readLines(lines []string) ([][]int64, []int64) {
	totals := make([]int64, 0)
	numberGroups := make([][]int64, 0)
	for _, line := range lines {
		splitTotalFromNumbers := strings.Split(line, ":")
		parsedTotal, err := strconv.ParseInt(splitTotalFromNumbers[0], 0, 64)
		totals = append(totals, parsedTotal)
		if err != nil {
			panic(err)
		}
		numbers := strings.Split(splitTotalFromNumbers[1], " ")
		parsedNumbers := make([]int64, 0)
		for _, number := range numbers {
			trimmedNumber := strings.TrimSpace(number)
			if len(trimmedNumber) > 0 {
				parseNumber, err := strconv.ParseInt(trimmedNumber, 0, 64)
				if err != nil {
					panic(err)
				}
				parsedNumbers = append(parsedNumbers, parseNumber)
			}
		}
		numberGroups = append(numberGroups, parsedNumbers)
	}
	return numberGroups, totals
}
