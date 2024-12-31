package day08

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"strings"
)

const Title = "Day 8: Resonant Collinearity"

var lines = utils.ReadFile("./data/day08/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

type Antenna struct {
	X int
	Y int
}

func PartOne() int {
	mapped, area := readLines(lines)
	list := make(map[[2]int]bool)
	for _, arr := range mapped {
		for index, _ := range arr {
			getAntiNodes(index, arr, list, area)
		}
	}
	return len(list)
}

func PartTwo() int {
	mapped, area := readLines(lines)
	list := make(map[[2]int]bool)
	for _, arr := range mapped {
		for index, _ := range arr {
			getHarmonicAntiNodes(index, arr, list, area)
		}
	}
	return len(list)
}

func getAntiNodes(index int, arr []Antenna, list map[[2]int]bool, lines []string) {
	antenna := arr[index]
	for currIdx, currAntenna := range arr {
		if currIdx == index {
			continue
		}
		distAntiX := (currAntenna.X - antenna.X) * 2
		distAntiY := (currAntenna.Y - antenna.Y) * 2
		antiX := antenna.X + distAntiX
		antiY := antenna.Y + distAntiY
		if antiY < 0 || antiY >= len(lines) {
			continue
		}
		if antiX < 0 || antiX >= len(lines[antiY]) {
			continue
		}
		list[[2]int{antiY, antiX}] = true
	}
}

func getHarmonicAntiNodes(index int, arr []Antenna, list map[[2]int]bool, lines []string) {
	antenna := arr[index]
	for currIdx, currAntenna := range arr {
		if currIdx == index {
			continue
		}
		distAntiX := currAntenna.X - antenna.X
		distAntiY := currAntenna.Y - antenna.Y
		antiX := antenna.X + distAntiX
		antiY := antenna.Y + distAntiY
		for {
			if antiY < 0 || antiY >= len(lines) {
				break
			}
			if antiX < 0 || antiX >= len(lines[antiY]) {
				break
			}
			list[[2]int{antiY, antiX}] = true
			antiX = antiX + distAntiX
			antiY = antiY + distAntiY
		}
	}
}

func readLines(lines []string) (map[string][]Antenna, []string) {
	area := make([]string, 0)
	mapped := make(map[string][]Antenna)
	for y, line := range lines {
		area = append(area, strings.TrimSpace(line))
		for x, char := range line {
			if !validateRune(char) {
				continue
			}
			stringified := string(char)
			_, ok := mapped[stringified]
			if !ok {
				mapped[stringified] = make([]Antenna, 0)
			}
			mapped[stringified] = append(mapped[stringified], Antenna{
				X: x,
				Y: y,
			})
		}
	}
	return mapped, area
}

func validateRune(char rune) bool {
	ascii := int(char)

	//digit
	if ascii > 47 && ascii < 58 {
		return true
	}

	// capital letters
	if ascii > 64 && ascii < 91 {
		return true
	}

	// minor letters
	if ascii > 96 && ascii < 123 {
		return true
	}

	return false
}
