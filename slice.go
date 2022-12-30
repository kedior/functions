package functions

func Filter[T any](fn func(T) bool, args ...T) []T {
	var t []T
	for _, arg := range args {
		if fn(arg) {
			t = append(t, arg)
		}
	}
	return t
}

func Map[A, B any](fn func(A) B, args ...A) []B {
	b := make([]B, len(args))
	for i, arg := range args {
		b[i] = fn(arg)
	}
	return b
}

func Reduce[T, R any](r R, fn func(R, T) R, args ...T) R {
	for _, arg := range args {
		r = fn(r, arg)
	}
	return r
}

func ReducePure[T any](fn func(T, T) T, args ...T) T {
	if len(args) == 0 {
		panic("args must not be empty")
	}
	return Reduce(args[0], fn, args[1:]...)
}

func Divide[T any](fn func(T) bool, args ...T) ([]T, []T) {
	var t, f []T
	for _, arg := range args {
		if fn(arg) {
			t = append(t, arg)
		} else {
			f = append(f, arg)
		}
	}
	return t, f
}

func Every[T any](fn func(T) bool, args ...T) bool {
	//return Reduce(true, func(b bool, t T) bool { return b && fn(t) }, args...)
	for _, arg := range args {
		if !fn(arg) {
			return false
		}
	}
	return true
}

func Some[T any](fn func(T) bool, args ...T) bool {
	//return Reduce(false, func(b bool, t T) bool { return b || fn(t) }, args...)
	for _, arg := range args {
		if fn(arg) {
			return true
		}
	}
	return false
}
