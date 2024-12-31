package day04

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
)

const Title = "Day 4: Ceres Search"

var lines = utils.ReadFile("./data/day04/input.txt")

var directions = [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

var nextLetterMap = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
}

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func checkAdjacents(column int, row int, directionColumn int, directionRow int, matrix [][]string) bool {
	curr := matrix[column][row]
	if curr == "S" {
		return true
	}
	adjX := column + directionColumn
	adjY := row + directionRow
	if adjX < 0 || adjX >= len(matrix) {
		return false
	}
	if adjY < 0 || adjY >= len(matrix[row]) {
		return false
	}
	adj := matrix[adjX][adjY]
	if nextLetterMap[curr] == adj {
		return checkAdjacents(adjX, adjY, directionColumn, directionRow, matrix)
	}
	return false
}

func PartOne() int {
	result := 0
	matrix := make([][]string, len(lines))
	for idx, line := range lines {
		for _, letter := range line {
			matrix[idx] = append(matrix[idx], string(letter))
		}
	}
	for column, line := range matrix {
		for row, element := range line {
			if element == "X" {
				for _, direction := range directions {
					foundXmas := checkAdjacents(column, row, direction[0], direction[1], matrix)
					if foundXmas {
						result++
					}
				}
			}
		}
	}
	return result
}

var crossDirections = [][]int{{-1, -1}, {-1, 1}}

func PartTwo() int {
	result := 0
	matrix := make([][]string, len(lines))
	for idx, line := range lines {
		for _, letter := range line {
			matrix[idx] = append(matrix[idx], string(letter))
		}
	}
	for column, line := range matrix {
		for row, element := range line {
			if element == "A" {
				sides := 0
				for _, direction := range crossDirections {
					adjCol := column + direction[0]
					adjRow := row + direction[1]
					invCol := column + (direction[0] * -1)
					invRow := row + (direction[1] * -1)
					if adjCol < 0 || adjCol >= len(matrix) {
						continue
					}
					if adjRow < 0 || adjRow >= len(matrix[adjCol]) {
						continue
					}
					if invCol < 0 || invCol >= len(matrix) {
						continue
					}
					if invRow < 0 || invRow >= len(matrix[invCol]) {
						continue
					}
					elementAtDir := matrix[adjCol][adjRow]
					elementAtInv := matrix[invCol][invRow]
					if elementAtDir == "S" && elementAtInv == "M" {
						sides++
					} else if elementAtInv == "S" && elementAtDir == "M" {
						sides++
					}
				}
				if sides == 2 {
					result++
				}
			}
		}
	}
	return result
}
