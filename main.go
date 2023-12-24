package main

import (
	"fmt"

	"github.com/nouhoum/dsa_practice/pkg/ds"
)

func main() {
	println("Hello, dsa!")
	a := ds.NewArrayList()
	a.Insert(8)
	a.Insert(32)
	a.Insert(54)
	fmt.Println(a)
	fmt.Println("==>inserting at 1 pos")
	a.InsertAt(1, 29)
	fmt.Println(a)
	fmt.Println("==>inserting at 2 pos")
	a.InsertAt(3, 7)
	fmt.Println(a)
	a0, _ := a.Get(0)
	fmt.Println("a[0]=", a0)
	a4, _ := a.Get(4)
	fmt.Println("a[4]=", a4)
}
