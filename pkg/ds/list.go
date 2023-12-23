package ds

type List interface {
	Len() int64
	Insert(e int64)
	InsertAt(i int64, e int64)
	String() string
}
