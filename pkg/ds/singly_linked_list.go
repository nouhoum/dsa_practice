package ds

import (
	"fmt"
	"strings"
)

type node struct {
	data int64
	next *node
}

type singlyLinkedList struct {
	head *node
}

// Get implements List.
func (*singlyLinkedList) Get(i int64) (int64, error) {
	panic("unimplemented")
}

// Insert implements List.
func (ll *singlyLinkedList) Insert(e int64) {
	newNode := &node{data: e}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	newNode.next = ll.head
	ll.head = newNode
}

// InsertAt implements List.
func (*singlyLinkedList) InsertAt(i int64, e int64) {
	panic("unimplemented")
}

// IsEmpty implements List.
func (ll *singlyLinkedList) IsEmpty() bool {
	return ll.head == nil
}

// Len implements List.
func (*singlyLinkedList) Len() int64 {
	panic("unimplemented")
}

func (*singlyLinkedList) RemoveAt(i int64) bool {
	panic("unimplemented")
}

func (ll *singlyLinkedList) String() string {
	if ll.head == nil {
		return "[]"
	}
	sb := strings.Builder{}
	sb.WriteString("[")
	curr := ll.head
	for curr != nil {
		sb.WriteString(fmt.Sprint(curr.data))
		curr = curr.next
		if curr != nil {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

// UpdateAt implements List.
func (*singlyLinkedList) UpdateAt(i int64, e int64) error {
	panic("unimplemented")
}

func NewSinglyLinkedList() List {
	return &singlyLinkedList{}
}
