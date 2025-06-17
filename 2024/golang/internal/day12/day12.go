package day12

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"strings"
)

const Title = "Day 12: Garden Groups"

var lines = utils.ReadFile("./data/day12/input.txt")

var directions = map[string][]int{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
}

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

// AREA = Points that touch outside
// PERIMETER = Number of elements
func PartOne() int {
	garden := readLines(lines)
    total := findRegions(garden)
	return total
}

func PartTwo() int {
	return 0
}

func findRegions(garden [][]string) int {
	seen := make([][]bool, len(garden))
    total := 0
	for idx := range garden {
		seen[idx] = make([]bool, len(garden[idx]))
	}
	for y, line := range garden {
		for x := range line {
			if seen[y][x] {
				continue
			}
			plotted, perimeter := bfs(y, x, garden, seen)
            total += len(plotted) * perimeter
		}
	}
    return total
}

func bfs(y, x int, garden [][]string, seen [][]bool) ([][]int, int) {
	queue := make([][]int, 0)
	queue = append(queue, []int{y, x})
	plotted := make([][]int, 0)
	perimeter := 0
	seen[y][x] = true
	for len(queue) > 0 {
		poppedQ, currentPos, _ := pop(queue)
		queue = poppedQ
		currentValue := garden[currentPos[0]][currentPos[1]]
		plotted = append(plotted, currentPos)
		for _, direction := range directions {
			nextPosY := currentPos[0] + direction[0]
			nextPosX := currentPos[1] + direction[1]
			if nextPosY < 0 || nextPosY >= len(garden) {
				perimeter++
				continue
			}
			if nextPosX < 0 || nextPosX >= len(garden[nextPosY]) {
				perimeter++
				continue
			}
			nextValue := garden[nextPosY][nextPosX]
			if nextValue != currentValue {
				perimeter++
				continue
			}
			if seen[nextPosY][nextPosX] {
				continue
			}
			queue = append(queue, []int{nextPosY, nextPosX})
			seen[nextPosY][nextPosX] = true
		}
	}
	return plotted, perimeter
}

func pop[T any](arr []T) ([]T, T, bool) {
	if len(arr) == 0 {
		var zero T
		return arr, zero, false
	}
	popped := arr[0]
	if len(arr) == 1 {
		arr = make([]T, 0)
	} else {
		arr = arr[1:]
	}
	return arr, popped, true

}

func readLines(lines []string) [][]string {
	garden := make([][]string, 0)
	for idx, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		garden = append(garden, make([]string, 0))
		for _, char := range trimmedLine {
			garden[idx] = append(garden[idx], string(char))
		}
	}
	return garden
}
