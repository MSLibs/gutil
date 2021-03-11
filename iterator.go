package gutil

type Iterator func() (item interface{}, ok bool)

type Query struct {
	Iterate func() Iterator
}

type Iterable interface {
	Iterate() Iterator
}
