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

func TestUpdateAt(t *testing.T) {
	cases := []struct {
		list               ds.List
		newElement         int64
		pos                int64
		expectedErr        error
		expectedListString string
	}{
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				return al
			}(),
			expectedListString: "[]",
			newElement:         123,
			expectedErr:        ds.ErrIndexOutOfBound,
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(4)
				al.UpdateAt(0, 33)
				return al
			}(),
			newElement:         33,
			pos:                0,
			expectedListString: "[ 33 ]",
			expectedErr:        nil,
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(4)
				al.Insert(55)
				al.Insert(1)
				return al
			}(),
			newElement:         42,
			pos:                1,
			expectedListString: "[ 4 42 1 ]",
			expectedErr:        nil,
		},
	}

	for _, test := range cases {
		err := test.list.UpdateAt(test.pos, test.newElement)
		if test.list.String() != test.expectedListString {
			t.Errorf("Expected string value to be %s, but got %s", test.expectedListString, test.list.String())
		}

		if err != test.expectedErr {
			t.Errorf("Expected string value to be %v, but got %v", test.expectedErr, err)
		}
	}
}

func TestRemoveAt(t *testing.T) {
	cases := []struct {
		list               ds.List
		pos                int64
		expectedFound      bool
		expectedListString string
	}{
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				return al
			}(),
			expectedListString: "[]",
			pos:                1,
			expectedFound:      false,
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(12)
				al.Insert(21)
				al.Insert(13)
				al.Insert(31)
				return al
			}(),
			expectedListString: "[ 12 21 31 ]",
			pos:                2,
			expectedFound:      true,
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(12)
				al.Insert(21)
				al.Insert(13)
				al.Insert(31)
				return al
			}(),
			expectedListString: "[ 12 21 13 ]",
			pos:                3,
			expectedFound:      true,
		},
		{
			list: func() ds.List {
				al := ds.NewArrayList()
				al.Insert(12)
				al.Insert(21)
				al.Insert(13)
				al.Insert(31)
				return al
			}(),
			expectedListString: "[ 12 21 13 31 ]",
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
