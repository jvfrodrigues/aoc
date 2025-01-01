package day09

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"strconv"
	"strings"
)

const Title = "Day 9: Disk Fragmenter"

var lines = utils.ReadFile("./data/day09/input.txt")

func Day() domain.AdventInterface {
	return domain.Advent[int]{
		Title:   Title,
		PartOne: PartOne,
		PartTwo: PartTwo,
	}
}

type DataBlock struct {
	ID   int
	Size int
}

func PartOne() int {
	dataBlocks := readLines(lines)
	compactData(dataBlocks)
	result := 0
	for idx, block := range dataBlocks {
		if block == -1 {
			break
		}
		result += idx * block
	}
	return result
}

func PartTwo() int {
	dataBlocks, emptyBlocks := readLinesStruct(lines)
	compactDefragedDataBlocks := compactDefragmentedDataBack(dataBlocks, emptyBlocks)
	result := 0
	totalIdx := 0
	for _, dataBlock := range compactDefragedDataBlocks {
		for _ = range dataBlock.Size {
			if dataBlock.ID != -1 {
				result += totalIdx * dataBlock.ID
			}
			totalIdx++
		}
	}
	return result
}

func compactData(dataBlocks []int) {
	idxLeft := 0
	idxRight := len(dataBlocks) - 1
	for idxLeft < idxRight {
		currLeft := dataBlocks[idxLeft]
		if currLeft != -1 {
			idxLeft++
			continue
		}
		currRight := dataBlocks[idxRight]
		if currRight == -1 {
			idxRight--
			continue
		}
		back := currLeft
		dataBlocks[idxLeft] = currRight
		dataBlocks[idxRight] = back
		idxLeft++
		idxRight--
	}
}

func compactDefragmentedData(dataBlocks []DataBlock, emptyBlocks map[int][]int) []DataBlock {
	idxRight := len(dataBlocks) - 1
	compactDefrag := make([]DataBlock, 0)
	compactDefrag = append(compactDefrag, dataBlocks...)
	for idxRight > 0 {
		currRight := compactDefrag[idxRight]
		if currRight.ID == -1 {
			idxRight--
			continue
		}
		currSize := currRight.Size
		for key, arr := range emptyBlocks {
			if key < currSize {
				continue
			}
			if len(arr) == 0 {
				continue
			}
			first := arr[0]
            if first >= idxRight {
                continue
            }
			emptyBlocks[key] = emptyBlocks[key][1:]
			sizeDiff := key - currRight.Size
			if sizeDiff == 0 {
				compactDefrag[first], compactDefrag[idxRight] = compactDefrag[idxRight], compactDefrag[first]
			} else {
				compactDefrag[first].Size = compactDefrag[idxRight].Size
				compactDefrag[first], compactDefrag[idxRight] = compactDefrag[idxRight], compactDefrag[first]
				idxRight++
				first++
				newElement := DataBlock{
					ID:   -1,
					Size: sizeDiff,
				}
				compactDefrag = append(compactDefrag[:first], append([]DataBlock{newElement}, compactDefrag[first:]...)...)
				emptyBlocks[sizeDiff] = append([]int{first}, emptyBlocks[sizeDiff]...)
			}
			curr := compactDefrag[idxRight]
			adjLeft := idxRight - 1
			adjRight := idxRight + 1
			if adjRight < len(compactDefrag) {
				adj := compactDefrag[adjRight]
				if adj.ID == -1 {
					curr.Size = curr.Size + adj.Size
					compactDefrag = append(compactDefrag[:idxRight], append([]DataBlock{curr}, compactDefrag[adjRight+1:]...)...)
				}
			}
			if adjLeft > -1 {
				adj := compactDefrag[adjLeft]
				if adj.ID == -1 {
					curr.Size = curr.Size + adj.Size
					compactDefrag = append(compactDefrag[:adjLeft], append([]DataBlock{curr}, compactDefrag[adjRight:]...)...)
				}
			}
			emptyBlocks = recheckPositions(compactDefrag)
			break
		}
		idxRight--
	}
	return compactDefrag
}

func compactDefragmentedDataBack(dataBlocks []DataBlock, emptyBlocks map[int][]int) []DataBlock {
	idxLeft := 0
	idxRight := len(dataBlocks) - 1
	compactDefrag := make([]DataBlock, 0)
	compactDefrag = append(compactDefrag, dataBlocks...)
	for idxRight > 0 {
		idxLeft = 0
		currRight := compactDefrag[idxRight]
		if currRight.ID == -1 {
			idxRight--
			continue
		}
		for idxLeft < idxRight {
			currLeft := compactDefrag[idxLeft]
			if currLeft.ID != -1 {
				idxLeft++
				continue
			}
			if currLeft.Size < currRight.Size {
				idxLeft++
				continue
			}
			sizeDiff := currLeft.Size - currRight.Size
			if sizeDiff == 0 {
				compactDefrag[idxLeft], compactDefrag[idxRight] = compactDefrag[idxRight], compactDefrag[idxLeft]
			} else {
				compactDefrag[idxLeft].Size = compactDefrag[idxRight].Size
				compactDefrag[idxLeft], compactDefrag[idxRight] = compactDefrag[idxRight], compactDefrag[idxLeft]
				idxLeft++
				idxRight++
				newElement := DataBlock{
					ID:   -1,
					Size: sizeDiff,
				}
				compactDefrag = append(compactDefrag[:idxLeft], append([]DataBlock{newElement}, compactDefrag[idxLeft:]...)...)
			}
			curr := compactDefrag[idxRight]
			adjLeft := idxRight - 1
			adjRight := idxRight + 1
			if adjRight < len(compactDefrag) {
				adj := compactDefrag[adjRight]
				if adj.ID == -1 {
					curr.Size = curr.Size + adj.Size
					compactDefrag = append(compactDefrag[:idxRight], append([]DataBlock{curr}, compactDefrag[adjRight+1:]...)...)
				}
			}
			if adjLeft > -1 {
				adj := compactDefrag[adjLeft]
				if adj.ID == -1 {
					curr.Size = curr.Size + adj.Size
					compactDefrag = append(compactDefrag[:adjLeft], append([]DataBlock{curr}, compactDefrag[adjRight:]...)...)
				}
			}
			break
		}
		idxRight--
	}
	return compactDefrag
}

func readLines(lines []string) []int {
	fileFlag := true
	dataBlocks := make([]int, 0)
	incID := 0
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		for _, char := range trimmedLine {
			fileSize, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				panic(err)
			}
			if fileFlag {
				fileFlag = false
				for _ = range fileSize {
					dataBlocks = append(dataBlocks, incID)
				}
				incID++
			} else {
				fileFlag = true
				for _ = range fileSize {
					dataBlocks = append(dataBlocks, -1)
				}
			}
		}
	}
	return dataBlocks
}

func readLinesStruct(lines []string) ([]DataBlock, map[int][]int) {
	fileFlag := false
	dataBlocks := make([]DataBlock, 0)
	availableBlocks := make(map[int][]int)
	incID := 0
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		for _, char := range trimmedLine {
			fileSize, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				panic(err)
			}
			fileFlag = !fileFlag
			if fileSize == 0 {
				continue
			}
			if fileFlag {
				dataBlocks = append(dataBlocks, DataBlock{
					ID:   incID,
					Size: int(fileSize),
				})
				incID++
			} else {
				dataBlocks = append(dataBlocks, DataBlock{
					ID:   -1,
					Size: int(fileSize),
				})
				_, ok := availableBlocks[int(fileSize)]
				if !ok {
					availableBlocks[int(fileSize)] = make([]int, 0)
				}
				availableBlocks[int(fileSize)] = append(availableBlocks[int(fileSize)], len(dataBlocks)-1)
			}
		}
	}
	return dataBlocks, availableBlocks
}

func recheckPositions(dataBlocks []DataBlock) map[int][]int {
	availableBlocks := make(map[int][]int)
	for idx, block := range dataBlocks {
		if block.ID != -1 {
			continue
		}
		blockSize := block.Size
		_, ok := availableBlocks[blockSize]
		if !ok {
			availableBlocks[blockSize] = make([]int, 0)
		}
		availableBlocks[blockSize] = append(availableBlocks[blockSize], idx)
	}
	return availableBlocks
}
