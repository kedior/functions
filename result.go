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

func Try2[T any](fn func() (error, T)) Result[T] {
	err, t := fn()
	return Result[T]{data: t, err: err}
}

func Try3[A, B any](fn func() (error, A, B)) Result2[A, B] {
	err, a, b := fn()
	return Result2[A, B]{a, b, err}
}

func RetRev[A, B any](fn func() (A, B)) func() (B, A) {
	a, b := fn()
	return func() (B, A) {
		return b, a
	}
}

func RetRevACB[A, B, C any](fn func() (A, B, C)) func() (A, C, B) {
	a, b, c := fn()
	return func() (A, C, B) {
		return a, c, b
	}
}

func RetRevBAC[A, B, C any](fn func() (A, B, C)) func() (B, A, C) {
	a, b, c := fn()
	return func() (B, A, C) {
		return b, a, c
	}
}

func RetRevBCA[A, B, C any](fn func() (A, B, C)) func() (B, C, A) {
	a, b, c := fn()
	return func() (B, C, A) {
		return b, c, a
	}
}

func RetRevCAB[A, B, C any](fn func() (A, B, C)) func() (C, A, B) {
	a, b, c := fn()
	return func() (C, A, B) {
		return c, a, b
	}
}

func RetRevCBA[A, B, C any](fn func() (A, B, C)) func() (C, B, A) {
	a, b, c := fn()
	return func() (C, B, A) {
		return c, b, a
	}
}
