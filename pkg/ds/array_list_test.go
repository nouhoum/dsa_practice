package ds_test

import (
	"math/rand"
	"testing"

	"github.com/nouhoum/dsa_practice/pkg/ds"
)

func TestInsert(t *testing.T) {
	a := ds.NewArrayList()

	if a.String() != "[]" {
		t.Errorf("Expected string value to be [], but got %s", a.String())
	}

	a.Insert(4)
	if a.String() != "[ 4 ]" {
		t.Errorf("Expected string value to be [ 4 ], but got %s", a.String())
	}

	a.Insert(55)
	if a.String() != "[ 4 55 ]" {
		t.Errorf("Expected string value to be [ 4 55 ], but got %s", a.String())
	}
}

func TestLen(t *testing.T) {
	a := ds.NewArrayList()

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

	if a.Len() != 0 {
		t.Errorf("Expected len value to be 0, but got %d", a.Len())
	}

	a.Insert(4)
	if a.Len() != 1 {
		t.Errorf("Expected len value to be 1, but got %d", a.Len())
	}

	a.Insert(55)
	a.Insert(14)
	if a.Len() != 3 {
		t.Errorf("Expected len value to be 3, but got %d", a.Len())
	}
}
