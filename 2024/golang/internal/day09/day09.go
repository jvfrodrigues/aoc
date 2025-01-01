package day09

import (
	"aoc/internal/domain"
	"aoc/internal/utils"
	"container/heap"
	"math"
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

type FileBlock struct {
	ID    int
	Size  int
	Start int
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
	fileBlocks, emptyHeaps := readLinesStruct(lines)
	compactDefragmentedData(fileBlocks, emptyHeaps)
	result := 0
	for _, fileBlock := range fileBlocks {
		if fileBlock.ID != -1 {
			idx := fileBlock.Start
			for i := 0; i < fileBlock.Size; i++ {
				result += fileBlock.ID * idx
				idx++
			}
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

func compactDefragmentedData(fileBlocks []FileBlock, emptyHeaps []domain.IntHeap) {
	idx := len(fileBlocks) - 1
	for idx > 0 {
		curr := fileBlocks[idx]
		if curr.ID == -1 {
			idx--
			continue
		}
		heapIndex, emptySpaceIndex := -1, math.MaxInt
		for i := curr.Size; i < 10; i++ {
			if len(emptyHeaps[i]) == 0 {
				continue
			}
			newEmptySpaceIndex := heap.Pop(&emptyHeaps[i]).(int)

			if newEmptySpaceIndex > curr.Start || newEmptySpaceIndex > emptySpaceIndex {
				heap.Push(&emptyHeaps[i], newEmptySpaceIndex)
				continue
			}

			if heapIndex != -1 {
				heap.Push(&emptyHeaps[heapIndex], emptySpaceIndex)
			}

			emptySpaceIndex = newEmptySpaceIndex
			heapIndex = i
		}

		if heapIndex == -1 {
			idx--
			continue
		}
		fileBlocks[idx].Start = emptySpaceIndex

		if heapIndex > curr.Size {
			newHeapIndex := heapIndex - curr.Size
			heap.Push(&emptyHeaps[newHeapIndex], emptySpaceIndex+curr.Size)
		}
		idx--
	}
}

func compactDefragmentedDataNotOptimized(dataBlocks []DataBlock, emptyHeaps [][]int) []DataBlock {
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

func readLinesStruct(lines []string) ([]FileBlock, []domain.IntHeap) {
	fileFlag := false
	fileBlocks := make([]FileBlock, 0)
	emptyHeaps := make([]domain.IntHeap, 10)
	for i, _ := range emptyHeaps {
		heap.Init(&emptyHeaps[i])
	}
	incID := 0
	totalIndex := 0
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
				fileBlocks = append(fileBlocks, FileBlock{
					ID:    incID,
					Size:  int(fileSize),
					Start: totalIndex,
				})
				incID++
			} else {
				heap.Push(&emptyHeaps[fileSize], totalIndex)
			}
			totalIndex += int(fileSize)
		}
	}
	return fileBlocks, emptyHeaps
}
