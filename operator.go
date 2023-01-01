package functions

func Substitute[X, Y, T any](fn func(X) Y, substitution func(T) X) func(T) Y {
	return func(t T) Y {
		return fn(substitution(t))
	}
}
