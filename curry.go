package functions

func Curry20[A, B any](fn func(A, B)) func(A) func(B) {
	return func(a A) func(B) {
		return func(b B) {
			fn(a, b)
		}
	}
}

func Curry30[A, B, C any](fn func(A, B, C)) func(A) func(B) func(C) {
	return func(a A) func(B) func(C) {
		return Curry20(func(b B, c C) {
			fn(a, b, c)
		})
	}
}

func Curry40[A, B, C, D any](fn func(A, B, C, D)) func(A) func(B) func(C) func(D) {
	return func(a A) func(B) func(C) func(D) {
		return Curry30(func(b B, c C, d D) {
			fn(a, b, c, d)
		})
	}
}

func Curry50[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A) func(B) func(C) func(D) func(E) {
	return func(a A) func(B) func(C) func(D) func(E) {
		return Curry40(func(b B, c C, d D, e E) {
			fn(a, b, c, d, e)
		})
	}
}

func Curry21[A, B, R any](fn func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return fn(a, b)
		}
	}
}

func Curry31[A, B, C, R any](fn func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return Curry21(func(b B, c C) R {
			return fn(a, b, c)
		})
	}
}

func Curry41[A, B, C, D, R any](fn func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return Curry31(func(b B, c C, d D) R {
			return fn(a, b, c, d)
		})
	}
}

func Curry51[A, B, C, D, E, R any](fn func(A, B, C, D, E) R) func(A) func(B) func(C) func(D) func(E) R {
	return func(a A) func(B) func(C) func(D) func(E) R {
		return Curry41(func(b B, c C, d D, e E) R {
			return fn(a, b, c, d, e)
		})
	}
}

func Curry22[A, B, X, Y any](fn func(A, B) (X, Y)) func(A) func(B) (X, Y) {
	return func(a A) func(B) (X, Y) {
		return func(b B) (X, Y) {
			return fn(a, b)
		}
	}
}

func Curry32[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(A) func(B) func(C) (X, Y) {
	return func(a A) func(B) func(C) (X, Y) {
		return Curry22(func(b B, c C) (X, Y) {
			return fn(a, b, c)
		})
	}
}

func Curry42[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A) func(B) func(C) func(D) (X, Y) {
	return func(a A) func(B) func(C) func(D) (X, Y) {
		return Curry32(func(b B, c C, d D) (X, Y) {
			return fn(a, b, c, d)
		})
	}
}

func Curry52[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A) func(B) func(C) func(D) func(E) (X, Y) {
	return func(a A) func(B) func(C) func(D) func(E) (X, Y) {
		return Curry42(func(b B, c C, d D, e E) (X, Y) {
			return fn(a, b, c, d, e)
		})
	}
}
