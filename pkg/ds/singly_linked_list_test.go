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

func TestSinglyLinkedList_Get(t *testing.T) {
	// Test case 1: Getting an element at a valid index
	list := ds.NewSinglyLinkedList()
	list.Insert(5)
	list.Insert(4)
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)
	expected := int64(3)
	result, err := list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	expected = int64(5)
	result, err = list.Get(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	expected = int64(1)
	result, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Getting an element at a negative index
	_, err = list.Get(-1)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Test case 3: Getting an element at an index greater than the length of the list
	_, err = list.Get(10)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestSinglyLinkedList_InsertAt(t *testing.T) {
	cases := []struct {
		pos          int64
		element      int64
		list         ds.List
		expectedList ds.List
	}{
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				return al
			}(),
			pos:     0,
			element: 23,
			expectedList: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(23)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(2)
				return al
			}(),
			pos:     0,
			element: 23,
			expectedList: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(23)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(2)
				return al
			}(),
			pos:     1,
			element: 34,
			expectedList: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(34)
				al.Insert(2)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
			pos:     2,
			element: 55,
			expectedList: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(55)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
			pos:     3,
			element: 55,
			expectedList: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(55)
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
		},
	}

	for _, test := range cases {
		test.list.InsertAt(test.pos, test.element)
		if test.list.String() != test.expectedList.String() {
			t.Errorf("Expected list to be %s, but got %s", test.expectedList, test.list)
		}
	}
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	cases := []struct {
		list               ds.List
		pos                int64
		expectedFound      bool
		expectedListString string
	}{
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				return al
			}(),
			expectedListString: "[]",
			pos:                1,
			expectedFound:      false,
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(31)
				al.Insert(13)
				al.Insert(21)
				al.Insert(12)
				return al
			}(),
			expectedListString: "[12 21 31]",
			pos:                2,
			expectedFound:      true,
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(31)
				al.Insert(13)
				al.Insert(21)
				al.Insert(12)
				return al
			}(),
			expectedListString: "[12 21 13]",
			pos:                3,
			expectedFound:      true,
		},
		{
			list: func() ds.List {
				al := ds.NewSinglyLinkedList()
				al.Insert(31)
				al.Insert(13)
				al.Insert(21)
				al.Insert(12)
				return al
			}(),
			expectedListString: "[12 21 13 31]",
			pos:                4,
			expectedFound:      false,
		},
	}

	for _, test := range cases {
		found := test.list.RemoveAt(test.pos)
		if test.list.String() != test.expectedListString {
			t.Errorf("Expected string value to be %s, but got %s", test.expectedListString, test.list.String())
		}

		if found != test.expectedFound {
			t.Errorf("Expected string value to be %v, but got %v", test.expectedFound, found)
		}
	}
}
