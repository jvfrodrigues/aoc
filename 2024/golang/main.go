package main

import (
	"aoc/internal/config"
	"aoc/internal/domain"
	"fmt"
	"os"
	"time"
)

func main() {
	var inputDay string
	if len(os.Args) == 2 {
		inputDay = os.Args[1]
	}

	fmt.Printf("\n%s %s %s\n", termCyan("--- Advent of Code:"), termPurple("2024"), termCyan("---"))

	var totalNanoSeconds int64
	for i, adventDay := range config.AdventDays {
		if len(inputDay) == 0 || inputDay == fmt.Sprintf("%d", i+1) {
			totalNanoSeconds += runWithTimer(adventDay)
		}
	}

	if len(inputDay) == 0 {
		fmt.Printf("\n%s %s %s\n\n", termCyan("--- Total run time:"), termPurple("%s", time.Duration(totalNanoSeconds).String()), termCyan("---"))
	}
}

func runWithTimer(adventDay domain.AdventInterface) int64 {
	fmt.Printf("\n%s\n", adventDay.GetTitle())

	partOneStartTime := time.Now()
	partOneResult := adventDay.RunPartOne()
	partOneDuration := time.Now().Sub(partOneStartTime)
	if partOneResult != -1 {
		fmt.Printf("  %s %s %s\n", termBlue("Part 1:"), termGreen("%d", partOneResult), termYellow("(%v)", partOneDuration))
	} else {
		fmt.Printf("  %s %s\n", termBlue("Part 1:"), termGrey("SKIPPED"))
		partOneDuration = 0
	}

	partTwoStartTime := time.Now()
	partTwoResult := adventDay.RunPartTwo()
	partTwoDuration := time.Now().Sub(partTwoStartTime)
	if partTwoResult != -1 {
		fmt.Printf("  %s %s %s\n", termBlue("Part 2:"), termGreen("%d", partTwoResult), termYellow("(%v)", partTwoDuration))
	} else {
		fmt.Printf("  %s %s\n", termBlue("Part 2:"), termGrey("SKIPPED"))
		partTwoDuration = 0
	}

	return partOneDuration.Nanoseconds() + partTwoDuration.Nanoseconds()
}

func termCyan(format string, a ...any) string {
	return "\033[36m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termGreen(format string, a ...any) string {
	return "\033[32m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termYellow(format string, a ...any) string {
	return "\033[33m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termBlue(format string, a ...any) string {
	return "\033[34m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termPurple(format string, a ...any) string {
	return "\033[35m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termGrey(format string, a ...any) string {
	return "\033[90m" + fmt.Sprintf(format, a...) + "\033[0m"
}
