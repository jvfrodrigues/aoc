package day05

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

const Title = "Day 5: Print Queue"

var lines = utils.ReadFile("./data/day05/input.txt")

type QElement struct {
	value int
	after map[int]bool
}

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func PartOne() int {
	middleResult := 0
	productionLines, mappedElements := readLines(lines)
	for _, line := range productionLines {
		isValid := checkValidity(line, mappedElements)
		if isValid {
			middleResult += line[int(math.Ceil(float64(len(line)/2)))]
		}
	}
	return middleResult
}

func PartTwo() int {
	middleResult := 0
	productionLines, mappedElements := readLines(lines)
	for _, line := range productionLines {
		isValid := checkValidity(line, mappedElements)
		if !isValid {
			reorderedQueue := reorderQueue(line, mappedElements)
			middleResult += reorderedQueue[int(math.Ceil(float64(len(line)/2)))]
		}
	}
	return middleResult
}

func reorderQueue(queue []int, mappedElements map[int]QElement) []int {
	copyArr := append([]int{}, queue...)
	slices.SortFunc(copyArr, func(a, b int) int {
		if a == b {
			return 0
		}
		elementA := mappedElements[a]
		_, ok := elementA.after[b]
		if !ok {
			return 1
		}
		return -1
	})
	return copyArr
}

func checkValidity(queue []int, mappedElements map[int]QElement) bool {
	queueMax := len(queue)
	for currIdx, element := range queue {
		qElement := mappedElements[element]
		idxRight := currIdx + 1
		for {
			if idxRight < queueMax {
				elementAfter := queue[idxRight]
				_, ok := qElement.after[elementAfter]
				if !ok {
					return false
				}
				idxRight++
			}
			if idxRight >= queueMax {
				break
			}
		}
	}
	return true
}

func readLines(lines []string) ([][]int, map[int]QElement) {
	gotOrders := false
	mappedElements := make(map[int]QElement)
	queues := make([][]int, 0)
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			gotOrders = true
			continue
		}
		if !gotOrders {
			ordered := strings.Split(trimmedLine, "|")
			parsedOrderedBefore, err := strconv.ParseInt(ordered[0], 0, 32)
			if err != nil {
				panic(err)
			}
			parsedOrderedAfter, err := strconv.ParseInt(ordered[1], 0, 32)
			if err != nil {
				panic(err)
			}
			_, ok := mappedElements[int(parsedOrderedBefore)]
			if !ok {
				mappedElements[int(parsedOrderedBefore)] = QElement{
					value: int(parsedOrderedBefore),
					after: make(map[int]bool),
				}
			}
			mappedElements[int(parsedOrderedBefore)].after[int(parsedOrderedAfter)] = true
		} else {
			lineProd := strings.Split(trimmedLine, ",")
			queuedLine := make([]int, 0)
			for _, element := range lineProd {
				parsedElement, err := strconv.ParseInt(element, 0, 32)
				if err != nil {
					panic(err)
				}
				queuedLine = append(queuedLine, int(parsedElement))
			}
			queues = append(queues, queuedLine)
		}
	}
	return queues, mappedElements
}
