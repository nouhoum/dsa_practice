package main

import (
	"fmt"

	"github.com/nouhoum/dsa_practice/pkg/ds"
)

func main() {
	println("Hello, dsa!")
	a := ds.NewArrayList()
	a.Insert(8)
	a.Insert(21)
	a.Insert(32)
	a.Insert(32)
	fmt.Println(a)
}
