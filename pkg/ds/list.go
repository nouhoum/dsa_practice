package ds

import "errors"

var (
	ErrIndexOutOfBound = errors.New("index out of bounds")
)

type List interface {
	Len() int64
	Insert(e int64)
	InsertAt(i int64, e int64)
	UpdateAt(i int64, e int64) error
	Get(i int64) (int64, error)
	String() string
}
