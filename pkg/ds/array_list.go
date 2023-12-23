package ds

import (
	"fmt"
	"strings"
)

const (
	INITIAL_CAPACITY = 3
)

type arrayList struct {
	data     []int64
	end      int64
	capacity int64
}

func NewArrayList() List {
	return &arrayList{
		data:     make([]int64, INITIAL_CAPACITY),
		end:      -1,
		capacity: INITIAL_CAPACITY,
	}
}

// Insert inserts a new element into the array list.
// It takes an int64 value as a parameter and does not return anything.
func (a *arrayList) Insert(e int64) {
	if a.end == a.capacity-1 {
		a.capacity = 2 * a.capacity
		newData := make([]int64, a.capacity)
		var i int64
		for ; i <= a.end; i++ {
			newData[i] = a.data[i]
		}
		a.data = newData
	}

	a.end++
	a.data[a.end] = e
}

// InsertAt implements List.
func (*arrayList) InsertAt(i int64, e int64) {
	panic("unimplemented")
}

// Len implements List.
func (*arrayList) Len() int64 {
	panic("unimplemented")
}

func (a *arrayList) String() string {
	if a.end == -1 {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[")
	var i int64
	for ; i <= a.end; i++ {
		sb.WriteString(fmt.Sprintf(" %d", a.data[i]))
	}
	sb.WriteString(" ]")
	return sb.String()
}

// List is empty: size = 0
// insert
// insertAt
// remove
// count
// read or modify element at position i
// specify list data type
// max size (capacity)
// int end = -1
