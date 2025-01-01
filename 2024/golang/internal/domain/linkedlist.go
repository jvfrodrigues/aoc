package domain

import "fmt"

type LinkedList[T any] struct {
	head   *LinkedElement[T]
	end    *LinkedElement[T]
	Length int
}

type LinkedElement[T any] struct {
	Value  T
	before *LinkedElement[T]
	after  *LinkedElement[T]
}

func (l *LinkedList[T]) Push(value T) {
	newElement := &LinkedElement[T]{
		Value: value,
	}
	currEnd := l.end
	if currEnd != nil {
		currEnd.after = newElement
		newElement.before = currEnd
	} else {
		l.head = newElement
	}
	l.end = newElement
	l.Length++
}

func (l *LinkedList[T]) Enqueue(value T) {
	newElement := &LinkedElement[T]{
		Value: value,
	}
	currHead := l.head
	if currHead != nil {
		currHead.before = newElement
		newElement.after = currHead
	} else {
		l.end = newElement
	}
	l.head = newElement
	l.Length++
}

func (l *LinkedList[T]) Add(value T, pos int) {
	if pos < 0 || pos >= l.Length {
		return
	}
	newElement := &LinkedElement[T]{
		Value: value,
	}
	if pos == 0 {
		l.Enqueue(value)
		return
	} else if pos == l.Length-1 {
		l.Push(value)
		return
	} else {
		currAtPos := l.Get(pos)
		newElement.before = currAtPos.before
		currAtPos.before.after = newElement
		newElement.after = currAtPos
		currAtPos.before = newElement
	}
	l.Length++
}

func (l *LinkedList[T]) Remove(pos int) {
	if pos < 0 || pos >= l.Length {
		return
	}
	fmt.Println("length bef", l.Length)
	currAtPos := l.Get(pos)
	if currAtPos.before != nil {
		fmt.Println("before")
		currAtPos.before.after = currAtPos.after
	}
	if currAtPos.after != nil {
		fmt.Println("after")
		currAtPos.after.before = currAtPos.before
	}
	if pos == 0 {
		fmt.Println("head")
		l.head = currAtPos.after
	}
	if pos == l.Length-1 {
		fmt.Println("end")
		l.end = currAtPos.before
	}
	l.Length--
	fmt.Println("length aft", l.Length)
}

func (l *LinkedList[T]) Get(pos int) *LinkedElement[T] {
	if pos < 0 || pos >= l.Length {
		return nil
	}

	idx := 0
	curr := l.head
	if pos > (l.Length / 2) {
		idx = l.Length - 1
		curr = l.end
		for idx > pos {
			curr = curr.before
			idx--
		}
	} else {
		for idx < pos {
			curr = curr.after
			idx++
		}
	}
	return curr
}
