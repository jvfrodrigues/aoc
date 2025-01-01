package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Push(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	assert.Equal(t, 3, ll.Length, "Length should be 3 after 3 pushes")
	assert.Equal(t, 3, ll.end.Value, "End value should be 3")
	assert.Equal(t, 1, ll.head.Value, "Head value should be 1")
}

func TestLinkedList_Enqueue(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Enqueue(1)
	ll.Enqueue(2)
	ll.Enqueue(3)

	assert.Equal(t, 3, ll.Length, "Length should be 3 after 3 enqueues")
	assert.Equal(t, 1, ll.end.Value, "End value should be 1")
	assert.Equal(t, 3, ll.head.Value, "Head value should be 3")
}

func TestLinkedList_Get(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	assert.Equal(t, 1, ll.Get(0).Value, "Element at position 0 should be 1")
	assert.Equal(t, 2, ll.Get(1).Value, "Element at position 1 should be 2")
	assert.Equal(t, 3, ll.Get(2).Value, "Element at position 2 should be 3")
	assert.Nil(t, ll.Get(3), "Element at position 3 should be nil")
}

func TestLinkedList_Add(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	ll.Add(4, 1)

	assert.Equal(t, 4, ll.Length, "Length should be 4 after adding an element")
	assert.Equal(t, 4, ll.Get(1).Value, "Element at position 1 should be 4")
	assert.Equal(t, 2, ll.Get(2).Value, "Element at position 2 should be 2")
}

func TestLinkedList_Remove(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Push(1)
	ll.Push(2)
	ll.Push(3)

	ll.Remove(1)

	assert.Equal(t, 2, ll.Length, "Length should be 2 after removing an element")
	assert.Equal(t, 3, ll.Get(1).Value, "Element at position 1 should now be 3")
	assert.Nil(t, ll.Get(2), "Element at position 2 should be nil")
}

func TestLinkedList_Integration(t *testing.T) {
	ll := &LinkedList[int]{}

	ll.Enqueue(1) // Head: 1
	ll.Push(2)    // Head: 1, End: 2
	ll.Add(3, 1)  // Head: 1, End: 2, Position 1: 3

	ll.Remove(0) // Remove head (1)

	assert.Equal(t, 2, ll.Length, "Length should be 2 after removing head")
	assert.Equal(t, 2, ll.head.Value, "Head should now be 2")
	assert.Equal(t, 3, ll.end.Value, "End should now be 3")
}
