package config

import (
	"aoc/internal/day01"
	"aoc/internal/day02"
	"aoc/internal/day03"
	"aoc/internal/day04"
	"aoc/internal/day05"
	"aoc/internal/day06"
	"aoc/internal/day07"
	"aoc/internal/day08"
	"aoc/internal/day09"
	"aoc/internal/domain"
)

var AdventDays = []domain.AdventInterface{
	day01.Day(),
	day02.Day(),
	day03.Day(),
	day04.Day(),
	day05.Day(),
	day06.Day(),
	day07.Day(),
	day08.Day(),
	day09.Day(),
}
