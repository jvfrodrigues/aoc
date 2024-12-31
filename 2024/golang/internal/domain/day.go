package domain

type AdventInterface interface {
	GetTitle() string
	RunPartOne() any
	RunPartTwo() any
}

type Advent[T int | int64] struct {
	Title   string
	PartOne func() T
	PartTwo func() T
}

func (a Advent[T]) GetTitle() string {
	return a.Title
}

func (a Advent[T]) RunPartOne() any {
	return a.PartOne()
}

func (a Advent[T]) RunPartTwo() any {
	return a.PartTwo()
}
