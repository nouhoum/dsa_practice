package ds_test

import (
	"math/rand"
	"testing"

	"github.com/nouhoum/dsa_practice/pkg/ds"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		list     ds.List
		expected string
	}{
		{
			list: func() ds.List {
				return ds.NewArrayList()
			}(),
			expected: "[]",
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(4)
				return al
			}(),
			expected: "[ 4 ]",
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(4)
				al.Insert(55)
				return al
			}(),
			expected: "[ 4 55 ]",
		},
	}

	for _, test := range cases {
		if test.list.String() != test.expected {
			t.Errorf("Expected string value to be %s, but got %s", test.expected, test.list.String())
		}
	}
}

func TestInsertAt(t *testing.T) {
	cases := []struct {
		pos          int64
		element      int64
		list         ds.List
		expectedList ds.List
	}{
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				return al
			}(),
			pos:     0,
			element: 23,
			expectedList: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(23)
				al.Insert(1)
				al.Insert(2)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				return al
			}(),
			pos:     1,
			element: 34,
			expectedList: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(34)
				al.Insert(2)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
			pos:     2,
			element: 55,
			expectedList: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(55)
				al.Insert(3)
				return al
			}(),
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				return al
			}(),
			pos:     3,
			element: 55,
			expectedList: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(1)
				al.Insert(2)
				al.Insert(3)
				al.Insert(55)
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

func TestLen(t *testing.T) {
	cases := []struct {
		al          ds.List
		expectedLen int64
	}{
		{
			expectedLen: 0,
			al: func() ds.List {
				return ds.NewArrayList()
			}(),
		},
		{
			expectedLen: 2,
			al: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(rand.Int63())
				al.Insert(rand.Int63())
				return al
			}(),
		},
		{
			expectedLen: 5,
			al: func() ds.List {
				al := ds.NewArrayList()
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
