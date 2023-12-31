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
		copy(newData, al.data)
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
		copy(newData, al.data)
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

func (al *arrayList) Get(i int64) (int64, error) {
	if i > al.end {
		return 0, ErrIndexOutOfBound
	}

	return al.data[i], nil
}

func (al *arrayList) UpdateAt(i int64, e int64) error {
	if i > al.end {
		return ErrIndexOutOfBound
	}

	al.data[i] = e
	return nil
}

func (al *arrayList) Len() int64 {
	return al.end + 1
}

// IsEmpty implements List.
func (al *arrayList) IsEmpty() bool {
	return al.end == -1
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

// RemoveAt removes an element at the specified index in the arrayList.
//
// Parameters:
// - i: the index of the element to be removed.
//
// Returns:
// - bool: true if the element was successfully removed, false otherwise.
func (al *arrayList) RemoveAt(i int64) bool {
	if i > al.end || al.end == -1 {
		return false
	}

	for j := i; j <= al.end; j++ {
		al.data[i] = al.data[i+1]
	}
	al.end--
	return true
}

// TODO:
// specify list data type
