package functions

type Result[T any] struct {
	data T
	err  error
}

type Result2[A, B any] struct {
	a   A
	b   B
	err error
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.data
}

func (r Result2[A, B]) Unwrap() (A, B) {
	if r.err != nil {
		panic(r.err)
	}
	return r.a, r.b
}

func (r Result[T]) Or(defaultV T) T {
	if r.err != nil {
		return defaultV
	}
	return r.data
}

func (r Result2[A, B]) Or(defaultA A, defaultB B) (A, B) {
	if r.err != nil {
		return defaultA, defaultB
	}
	return r.a, r.b
}

func Try2[T any](val T, err error) Result[T] {
	return Result[T]{data: val, err: err}
}

func Try3[A, B any](a A, b B, err error) Result2[A, B] {
	return Result2[A, B]{a, b, err}
}
