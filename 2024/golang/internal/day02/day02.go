package day02

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"math"
	"strconv"
	"strings"
)

const Title = "Day 2: Red-Nosed Reports"

var lines = utils.ReadFile("./data/day02/inputgui.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

func PartOne() int {
	safe := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		levelsNum := make([]int, len(levels))
		for idx, level := range levels {
			levelParsed, err := strconv.ParseInt(strings.TrimSpace(level), 0, 32)
			if err != nil {
				panic(err)
			}
			levelsNum[idx] = int(levelParsed)
		}
		isSafe := checkReport(levelsNum)

		if isSafe {
			safe++
		}
	}
    return safe
}

func PartTwo() int {
	safe := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		levelsNum := make([]int, len(levels))
		for idx, level := range levels {
			levelParsed, err := strconv.ParseInt(strings.TrimSpace(level), 0, 32)
			if err != nil {
				panic(err)
			}
			levelsNum[idx] = int(levelParsed)
		}
		isSafe := checkReport(levelsNum)

        // PART 2 DIFFERENCE
		if !isSafe {
			for idx, _ := range levelsNum {
                copyArr := append([]int{}, levelsNum...)
                alteredLevels := utils.SliceRemove(copyArr, idx)
				isAlteredSafe := checkReport(alteredLevels)
				if isAlteredSafe {
					isSafe = true
				}
			}
		}

		if isSafe {
			safe++
		}
	}
    return safe
}

func checkReport(levelsNum []int) bool {
	direction := 0
	for i := 0; i < len(levelsNum)-1; i++ {
		level1 := levelsNum[i]
		level2 := levelsNum[i+1]

		difference := level2 - level1
		absDifference := math.Abs(float64(difference))
		if absDifference < 1 || absDifference > 3 {
			return false
		}

		if direction == 0 {
			if difference < 0 {
				direction = -1
			} else {
				direction = 1
			}
		}

		if difference < 0 && direction > 0 {
			return false
		} else if difference > 0 && direction < 0 {
			return false
		}
	}
	return true
}
