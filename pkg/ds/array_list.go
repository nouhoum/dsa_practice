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

func (al *arrayList) InsertAt(pos int64, e int64) {
	if pos > al.end {
		al.Insert(e)
		return
	}

	var newData []int64
	if al.end+1 == al.capacity {
		al.capacity = al.capacity * 2
		newData = make([]int64, al.capacity)

		for j := int64(0); j < pos; j++ {
			newData[j] = al.data[j]
		}
	} else {
		newData = al.data
	}

	for j := al.end; j >= pos; j-- {
		newData[j+1] = al.data[j]
	}

	newData[pos] = e
	al.end++
	al.data = newData
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

// TODO:
// remove
// read or modify element at position i
// specify list data type
