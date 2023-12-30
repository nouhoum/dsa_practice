package ds_test

import (
	"math/rand"
	"testing"

	"github.com/nouhoum/dsa_practice/pkg/ds"
)

func TestSinglyLinkedList_Len(t *testing.T) {
	cases := []struct {
		al          ds.List
		expectedLen int64
	}{
		{
			expectedLen: 0,
			al: func() ds.List {
				return ds.NewSinglyLinkedList()
			}(),
		},
		{
			expectedLen: 2,
			al: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				return al
			}(),
		},
		{
			expectedLen: 5,
			al: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				return al
			}(),
		},
	}

	for _, test := range cases {
		if test.expectedLen != test.al.Len() {
			t.Errorf("Expected len value to be %d, but got %d", test.expectedLen, test.al.Len())
		}
	}
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	// Test case 1: head is nil, should return true
	ll1 := ds.NewSinglyLinkedList()
	if !ll1.IsEmpty() {
		t.Error("Expected IsEmpty() to return true, but it returned false")
	}

	// Test case 2: head is not nil, should return false
	ll2 := ds.NewSinglyLinkedList()
	ll2.Insert(1)
	if ll2.IsEmpty() {
		t.Error("Expected IsEmpty() to return false, but it returned true")
	}
}

func TestSinglyLinkedList_String(t *testing.T) {
	// Test case 1: Empty list
	ll := ds.NewSinglyLinkedList()
	expected := "[]"
	if result := ll.String(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 2: Non-empty list
	ll = ds.NewSinglyLinkedList()
	ll.Insert(1)

	expected = "[1]"
	if result := ll.String(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	ll.Insert(2)
	expected = "[2 1]"
	if result := ll.String(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	ll.Insert(3)
	expected = "[3 2 1]"
	if result := ll.String(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// func TestSinglyLinkedList_Insert(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		list ds.List
// 		e    int64
// 	}{
// 		{
// 			name: "Inserting into an empty list",
// 			list: ds.NewSinglyLinkedList(),
// 			e:    1,
// 		},
// 		{
// 			name: "Inserting into a non-empty list",
// 			list: singlyLinkedList{
// 				head: &node{value: 1},
// 			},
// 			e: 2,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.list.Insert(tt.e)
// 			// Add assertions here to check if the element was inserted correctly.
// 		})
// 	}
// }

// func TestSinglyLinkedList_Get(t *testing.T) {
// 	// Test case 1: Getting an element at a valid index
// 	// list := &singlyLinkedList{1, 2, 3, 4, 5}
// 	list := ds.NewSinglyLinkedList()
// 	expected := int64(3)
// 	result, err := list.Get(2)
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}
// 	if result != expected {
// 		t.Errorf("Expected %v, got %v", expected, result)
// 	}

// 	// Test case 2: Getting an element at a negative index
// 	// list = &singlyLinkedList{1, 2, 3, 4, 5}
// 	list = ds.NewSinglyLinkedList()
// 	_, err = list.Get(-1)
// 	if err == nil {
// 		t.Errorf("Expected an error, got nil")
// 	}

// 	// Test case 3: Getting an element at an index greater than the length of the list
// 	// list = &singlyLinkedList{1, 2, 3, 4, 5}
// 	list = ds.NewSinglyLinkedList()
// 	_, err = list.Get(10)
// 	if err == nil {
// 		t.Errorf("Expected an error, got nil")
// 	}
// }
