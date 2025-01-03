package day10

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"strconv"
	"strings"
	"sync"
)

const Title = "Day 10: Hoof It"

var lines = utils.ReadFile("./data/day10/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

const (
	trailHead int = 0
	trailEnd  int = 9
)

var directions = map[string][]int{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
}

func PartOne() int {
	matrix := readLines(lines)
	result := 0
	c := make(chan bool, len(matrix)*len(matrix))
	wg := sync.WaitGroup{}
	for Y, line := range matrix {
		for X, element := range line {
			if element != trailHead {
				continue
			}
			seen := make(map[[2]int]bool)
			wg.Add(1)
			go func() {
				defer wg.Done()
				findTrailsFromHead(Y, X, element, matrix, seen, c)
			}()
		}
	}
	wg.Wait()
	close(c)
	for _ = range c {
		result++
	}
	return result
}

func PartTwo() int {
	matrix := readLines(lines)
	result := 0
	c := make(chan bool)
	wg := sync.WaitGroup{}
	for Y, line := range matrix {
		for X, element := range line {
			if element != trailHead {
				continue
			}
			seen := make(map[[2]int]bool)
			wg.Add(1)
			go func() {
				defer wg.Done()
				findTrailsFromHeadP2(Y, X, element, matrix, seen, c)
			}()
		}
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for _ = range c {
		result++
	}
	return result
}

func findTrailsFromHead(y, x, currValue int, matrix [][]int, seen map[[2]int]bool, c chan bool) {
	if currValue == trailEnd {
		c <- true
		return
	}
	for _, direction := range directions {
		adjPosY := y + direction[0]
		adjPosX := x + direction[1]
		if adjPosY < 0 || adjPosY >= len(matrix) {
			continue
		}
		if adjPosX < 0 || adjPosX >= len(matrix[adjPosY]) {
			continue
		}
		nextValue := matrix[adjPosY][adjPosX]
		if nextValue != currValue+1 {
			continue
		}
		_, ok := seen[[2]int{adjPosY, adjPosX}]
		if ok {
			continue
		}
		seen[[2]int{adjPosY, adjPosX}] = true
		findTrailsFromHead(adjPosY, adjPosX, nextValue, matrix, seen, c)
	}
}

func findTrailsFromHeadP2(y, x, currValue int, matrix [][]int, seen map[[2]int]bool, c chan bool) {
	if currValue == trailEnd {
		c <- true
		return
	}
	for _, direction := range directions {
		directionSeen := make(map[[2]int]bool)
		for key, value := range seen {
			directionSeen[key] = value
		}
		adjPosY := y + direction[0]
		adjPosX := x + direction[1]
		if adjPosY < 0 || adjPosY >= len(matrix) {
			continue
		}
		if adjPosX < 0 || adjPosX >= len(matrix[adjPosY]) {
			continue
		}
		nextValue := matrix[adjPosY][adjPosX]
		if nextValue != currValue+1 {
			continue
		}
		_, ok := directionSeen[[2]int{adjPosY, adjPosX}]
		if ok {
			continue
		}
		directionSeen[[2]int{adjPosY, adjPosX}] = true
		findTrailsFromHeadP2(adjPosY, adjPosX, nextValue, matrix, directionSeen, c)
	}
}

func readLines(lines []string) [][]int {
	matrix := make([][]int, 0)
	for Y, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		matrix = append(matrix, []int{})
		for _, char := range trimmedLine {
			parsedNumber, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				panic(err)
			}
			matrix[Y] = append(matrix[Y], int(parsedNumber))
		}
	}
	return matrix
}
