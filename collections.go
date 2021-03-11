package gutil

import "reflect"

// 按指定条件过滤数组

type ArraySource struct {
	source []interface{}
}

func From(source interface{}) Query {
	src := reflect.ValueOf(source)
	switch src.Kind() {
	case reflect.Array, reflect.Slice:
		len := src.Len()
		return Query{
			Iterate: func() Iterator {
				index := 0
				return func() (item interface{}, ok bool) {
					ok = index < len
					if ok {
						item = src.Index(index).Interface()
						index++
					}
					return
				}
			},
		}
	default:
		return FromIterable(source.(Iterable))
	}
}

func FromIterable(source Iterable) Query {
	return Query{
		Iterate: source.Iterate,
	}
}

func FromArray(source []interface{}) *ArraySource {
	return &ArraySource{source: source}
}

func (s ArraySource) Any(predicate func(interface{}) bool) bool {
	if len(s.source) == 0 {
		return false
	}
	for _, item := range s.source {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s ArraySource) All(predicate func(interface{}) bool) bool {
	if len(s.source) == 0 {
		return false
	}
	for _, item := range s.source {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (ArraySource) Where(predicate func(interface{}) bool) *Query {
	return &Query{}
}
