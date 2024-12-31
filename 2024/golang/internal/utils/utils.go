package utils

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lines := make([]string, 0)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		if len(line) > 0 {
			lines = append(lines, string(line))
		}
	}
	return lines
}

func SliceRemove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
