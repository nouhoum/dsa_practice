package ds_test

import (
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