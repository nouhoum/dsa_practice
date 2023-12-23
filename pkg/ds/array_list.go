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
func (al *arrayList) Insert(e int64) {
	if al.end == al.capacity-1 {
		al.capacity = 2 * al.capacity
		newData := make([]int64, al.capacity)
		var i int64
		for ; i <= al.end; i++ {
			newData[i] = al.data[i]
		}
		al.data = newData
	}

	al.end++
	al.data[al.end] = e
}

// InsertAt implements List.
func (*arrayList) InsertAt(i int64, e int64) {
	panic("unimplemented")
}

func (al *arrayList) Len() int64 {
	return al.end + 1
}

func (al *arrayList) String() string {
	if al.end == -1 {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[")
	var i int64
	for ; i <= al.end; i++ {
		sb.WriteString(fmt.Sprintf(" %d", al.data[i]))
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
