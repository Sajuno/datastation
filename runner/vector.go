package runner

import "math"

type Vector[T any] struct {
	data  []T
	index int
}

func (v *Vector[T]) Insert(i int, item T) {
	if len(v.data) == 0 {
		v.data = make([]T, 8)
	}

	if i >= len(v.data) {
		rest := make([]T, int(math.Floor(float64(len(v.data))*.75)))
		v.data = append(v.data, rest...)
	}

	v.data[i] = item
}

func (v *Vector[T]) Reset() {
	v.index = 0
}

func (v *Vector[T]) Append(item T) {
	v.Insert(v.index, item)
	v.index++
}

func (v *Vector[T]) List() []T {
	return v.data
}

func (v *Vector[T]) Index() int {
	return v.index
}
