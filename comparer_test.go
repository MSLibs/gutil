package gutil

import "testing"

func TestComparer(t *testing.T) {
	tests := []struct {
		x    interface{}
		y    interface{}
		want int
	}{
		{foo{v1: 100}, foo{v1: 200}, -1},
		{foo{v1: 100}, foo{v1: 100}, 0},
		{foo{v1: 200}, foo{v1: 100}, 1},
	}

	for _, e := range tests {
		if r := getInterfaceComparer(e.x)(e.x, e.y); r != e.want {
			t.Errorf("getComparer(%v)(%v,%v)=%v expected %v", e.x, e.x, e.y, r, e.want)
		}
	}
}

type foo struct {
	v1 int
	v2 bool
	v3 string
}

func (f foo) CompareTo(c Comparable) int {
	if ff, ok := c.(foo); ok {
		x, y := f.v1, ff.v1
		if x < y {
			return -1
		} else if x == y {
			return 0
		} else {
			return 1
		}
	}
	panic("undefined CompareTo")
}
