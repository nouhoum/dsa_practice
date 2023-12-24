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
