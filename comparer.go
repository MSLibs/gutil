package gutil

type comparer func(interface{}, interface{}) int

type Comparable interface {
	CompareTo(Comparable) int
}

func getInterfaceComparer(data interface{}) comparer {
	return func(x, y interface{}) int {
		a, b := x.(Comparable), y.(Comparable)
		return a.CompareTo(b)
	}
}
