package ds

import "errors"

var (
	ErrIndexOutOfBound = errors.New("index out of bounds")
)

type List interface {
	Insert(e int64)
	InsertAt(i int64, e int64)
	UpdateAt(i int64, e int64) error
	RemoveAt(i int64) bool
	Get(i int64) (int64, error)
	Len() int64
	IsEmpty() bool

	String() string
}
