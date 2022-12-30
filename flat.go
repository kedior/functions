package functions

func Fn1[A any](a A) func() A {
	return func() A {
		return a
	}
}

func Fn2[A, B any](a A, b B) func() (A, B) {
	return func() (A, B) {
		return a, b
	}
}

func Fn3[A, B, C any](a A, b B, c C) func() (A, B, C) {
	return func() (A, B, C) {
		return a, b, c
	}
}

func Fn4[A, B, C, D any](a A, b B, c C, d D) func() (A, B, C, D) {
	return func() (A, B, C, D) {
		return a, b, c, d
	}
}
func Fn5[A, B, C, D, E any](a A, b B, c C, d D, e E) func() (A, B, C, D, E) {
	return func() (A, B, C, D, E) {
		return a, b, c, d, e
	}
}
