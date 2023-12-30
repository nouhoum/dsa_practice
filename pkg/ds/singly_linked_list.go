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

func NewSinglyLinkedList() List {
	return &singlyLinkedList{}
}

func (ll *singlyLinkedList) Get(i int64) (int64, error) {
	if ll.head == nil {
		return 0, ErrIndexOutOfBound
	}
	var pos int64
	curr := ll.head
	for curr != nil {
		if pos == i {
			return curr.data, nil
		}
		curr = curr.next
		pos++
	}

	return 0, ErrIndexOutOfBound
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
func (ll *singlyLinkedList) InsertAt(i int64, e int64) {
	newNode := &node{data: e}
	if i == 0 { // We insert at the begining
		if ll.head == nil {
			ll.head = &node{data: e}
			return
		}

		newNode.next = ll.head
		ll.head = newNode
	}

	var pos int64
	curr := ll.head
	for curr != nil {
		if pos == i-1 { // we should insert after current
			newNode.next = curr.next
			curr.next = newNode
			return
		}
		curr = curr.next
		pos++
	}
}

func (ll *singlyLinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *singlyLinkedList) Len() int64 {
	var count int64
	curr := ll.head
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (ll *singlyLinkedList) RemoveAt(i int64) bool {
	if ll.head == nil {
		return false
	}

	curr := ll.head
	var pos int64
	for curr != nil {
		if pos == i-1 { // the next element should be removed
			next := curr.next
			if next == nil { // it's the end of the list, so no element to remove.
				return false
			}

			curr.next = next.next
			return true
		}

		pos++
		curr = curr.next
	}

	return false
}

func (ll *singlyLinkedList) UpdateAt(i int64, e int64) error {
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
