package binary

import (
	"github.com/pkg/errors"
	"reflect"
)

// Less is an interface to handle less func for type
type Less interface {
	Less() func(l, r interface{}) bool
}

// IntLess is standard Less func for int type
func IntLess() func(l, r interface{}) bool {
	return func(l, r interface{}) bool {
		lhs := l.(int)
		rhs := r.(int)
		return lhs < rhs
	}
}

// StringLess is standard Less func for string type
func StringLess() func(l, r interface{}) bool {
	return func(l, r interface{}) bool {
		lhs := l.(string)
		rhs := r.(string)
		return lhs < rhs
	}
}

// Search is a binary search function for different types
func Search(input, find interface{}, less func(l, r interface{}) bool) (int, error) {
	var n int
	if reflect.TypeOf(input).Kind() != reflect.Slice {
		return -1, errors.New("input is not a slice type")
	}

	v := reflect.ValueOf(input)
	n = v.Len() - 1

	for idx := 0; idx <= n; {
		m := (idx + n) / 2
		lhs := v.Index(m).Interface()
		rhs := find

		if less(lhs, rhs) {
			idx = m + 1
		} else if reflect.DeepEqual(lhs, rhs) {
			return m, nil
		} else {
			n = m - 1
		}
	}
	return -1, nil
}
