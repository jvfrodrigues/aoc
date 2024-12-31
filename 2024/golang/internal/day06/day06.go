package day06

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"slices"
)

const Title = "Day 6: Guard Gallivant"

var lines = utils.ReadFile("./data/day06/input.txt")

const (
	patrol = "^"
	wall   = "#"
)

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

func changeDirection(curr string) string {
	if curr == "up" {
		return "right"
	} else if curr == "right" {
		return "down"
	} else if curr == "down" {
		return "left"
	}
	return "up"
}

func PartOne() int {
	walkedInSpaces := 0
	area, startPoint, _ := readLines(lines)
	seen := make(map[[2]int]bool)
	followPath(startPoint[0], startPoint[1], area, seen, "up")
	walkedInSpaces = len(seen)
	return walkedInSpaces
}

func PartTwo() int {
	area, startPoint, obstacles := readLines(lines)
	for _, obstacleArr := range obstacles {
		slices.Sort(obstacleArr)
	}
	seen := make(map[[2]int]bool)
	possibleBlocks := waysToBlock(startPoint[0], startPoint[1], area, obstacles, "up", seen)
	return possibleBlocks
}

func readLines(lines []string) ([][]string, []int, map[string][]int) {
	area := make([][]string, 0)
	startPoint := make([]int, 0)
	obstacles := make(map[string][]int)
	for height, line := range lines {
		areaLine := make([]string, 0)
		for length, char := range line {
			if string(char) == patrol {
				startPoint = append(startPoint, height, length)
			} else if string(char) == wall {
				addObstacle(height, length, obstacles)
			}
			areaLine = append(areaLine, string(char))
		}
		area = append(area, areaLine)
	}
	return area, startPoint, obstacles
}

func addObstacle(height, length int, obstacles map[string][]int) {
	_, ok := obstacles["y"+string(height)]
	if !ok {
		obstacles["y"+string(height)] = make([]int, 0)
	}
	obstacles["y"+string(height)] = append(obstacles["y"+string(height)], length)
	_, ok = obstacles["x"+string(length)]
	if !ok {
		obstacles["x"+string(length)] = make([]int, 0)
	}
	obstacles["x"+string(length)] = append(obstacles["x"+string(length)], height)
}

func followPath(height, length int, area [][]string, seen map[[2]int]bool, direction string) {
	_, ok := seen[[2]int{height, length}]
	if !ok {
		seen[[2]int{height, length}] = true
	}
	nextDirection := directions[direction]
	newPosHeight := height + nextDirection[0]
	newPosLength := length + nextDirection[1]
	if newPosHeight < 0 || newPosHeight >= len(area) || newPosLength < 0 || newPosLength >= len(area[height]) {
		return
	}
	element := area[newPosHeight][newPosLength]
	if element == wall {
		direction = changeDirection(direction)
		followPath(height, length, area, seen, direction)
		return
	}
	followPath(newPosHeight, newPosLength, area, seen, direction)
	return
}

func waysToBlock(height, length int, area [][]string, obstacles map[string][]int, direction string, seen map[[2]int]bool) int {
	_, ok := seen[[2]int{height, length}]
	if !ok {
		seen[[2]int{height, length}] = true
	}
	count := 0
	nextDirection := directions[direction]
	newPosHeight := height + nextDirection[0]
	newPosLength := length + nextDirection[1]
	if newPosHeight < 0 || newPosHeight >= len(area) || newPosLength < 0 || newPosLength >= len(area[height]) {
		return count
	}
	element := area[newPosHeight][newPosLength]
	if element == wall {
		direction = changeDirection(direction)
		count += waysToBlock(height, length, area, obstacles, direction, seen)
		return count
	} else {
		canIt := checkItCanLoop(height, length, direction, area, obstacles, seen)
		if canIt {
			count++
		}
	}
	count += waysToBlock(newPosHeight, newPosLength, area, obstacles, direction, seen)
	return count
}

func checkItCanLoop(height, length int, direction string, area [][]string, obstacles map[string][]int, seen map[[2]int]bool) bool {
	nextDirection := changeDirection(direction)
	seenObstacles := make(map[[2]int]map[string]bool)

	possibleObstacleHeight := height + directions[direction][0]
	possibleObstacleLength := length + directions[direction][1]
	_, ok := seen[[2]int{possibleObstacleHeight, possibleObstacleLength}]
	if ok {
		return false
	}

	copyObstacles := make(map[string][]int)
	for key, value := range obstacles {
		copyObstacles[key] = make([]int, 0)
		copyObstacles[key] = append(copyObstacles[key], value...)
	}
	_, ok = copyObstacles["y"+string(possibleObstacleHeight)]
	if !ok {
		copyObstacles["y"+string(possibleObstacleHeight)] = make([]int, 0)
	}
	copyObstacles["y"+string(possibleObstacleHeight)] = append(copyObstacles["y"+string(possibleObstacleHeight)], possibleObstacleLength)
	slices.Sort(copyObstacles["y"+string(possibleObstacleHeight)])
	_, ok = copyObstacles["x"+string(possibleObstacleLength)]
	if !ok {
		copyObstacles["x"+string(possibleObstacleLength)] = make([]int, 0)
	}
	copyObstacles["x"+string(possibleObstacleLength)] = append(copyObstacles["x"+string(possibleObstacleLength)], possibleObstacleHeight)
	slices.Sort(copyObstacles["x"+string(possibleObstacleLength)])

	arr := [2]int{possibleObstacleHeight, possibleObstacleLength}
	seenObstacles[arr] = make(map[string]bool)
	seenObstacles[arr][direction] = true

	possiblePosHeight := height
	possiblePosLength := length
	for {
		closest := [2]int{-1, -1}
		if nextDirection == "up" || nextDirection == "down" {
			heights, ok := copyObstacles["x"+string(possiblePosLength)]
			if !ok {
				return false
			}
			if nextDirection == "up" {
				for _, height := range heights {
					if height < possiblePosHeight {
						closest[0] = height
						closest[1] = possiblePosLength
					} else {
						break
					}
				}
			} else {
				for _, height := range heights {
					if height > possiblePosHeight {
						closest[0] = height
						closest[1] = possiblePosLength
						break
					}
				}
			}
		} else {
			lengths, ok := copyObstacles["y"+string(possiblePosHeight)]
			if !ok {
				return false
			}
			if nextDirection == "left" {
				for _, length := range lengths {
					if length < possiblePosLength {
						closest[0] = possiblePosHeight
						closest[1] = length
					} else {
						break
					}
				}
			} else {
				for _, length := range lengths {
					if length > possiblePosLength {
						closest[0] = possiblePosHeight
						closest[1] = length
						break
					}
				}
			}
		}
		if closest[0] == -1 {
			return false
		}
		_, ok := seenObstacles[closest]
		if !ok {
			seenObstacles[closest] = make(map[string]bool)
		}
		_, ok = seenObstacles[closest][nextDirection]
		if ok {
			return true
		}
		seenObstacles[closest][nextDirection] = true
		possiblePosHeight = closest[0] - directions[nextDirection][0]
		possiblePosLength = closest[1] - directions[nextDirection][1]
		nextDirection = changeDirection(nextDirection)
	}
}
