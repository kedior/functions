package functions

func ArgRev0BA[A, B any](fn func(A, B)) func(B, A) {
	return func(b B, a A) {
		fn(a, b)
	}
}

func ArgRev1BA[A, B, X any](fn func(A, B) X) func(B, A) X {
	return func(b B, a A) X {
		return fn(a, b)
	}
}

func ArgRev2BA[A, B, X, Y any](fn func(A, B) (X, Y)) func(B, A) (X, Y) {
	return func(b B, a A) (X, Y) {
		return fn(a, b)
	}
}

func ArgRev3BA[A, B, X, Y, Z any](fn func(A, B) (X, Y, Z)) func(B, A) (X, Y, Z) {
	return func(b B, a A) (X, Y, Z) {
		return fn(a, b)
	}
}

func ArgRev0ACB[A, B, C any](fn func(A, B, C)) func(A, C, B) {
	return func(a A, c C, b B) {
		fn(a, b, c)
	}
}
func ArgRev0BAC[A, B, C any](fn func(A, B, C)) func(B, A, C) {
	return func(b B, a A, c C) {
		fn(a, b, c)
	}
}

func ArgRev0BCA[A, B, C any](fn func(A, B, C)) func(B, C, A) {
	return func(b B, c C, a A) {
		fn(a, b, c)
	}
}

func ArgRev0CAB[A, B, C any](fn func(A, B, C)) func(C, A, B) {
	return func(c C, a A, b B) {
		fn(a, b, c)
	}
}

func ArgRev0CBA[A, B, C any](fn func(A, B, C)) func(C, B, A) {
	return func(c C, b B, a A) {
		fn(a, b, c)
	}
}

func ArgRev1ACB[A, B, C, X any](fn func(A, B, C) X) func(A, C, B) X {
	return func(a A, c C, b B) X {
		return fn(a, b, c)
	}
}

func ArgRev1BAC[A, B, C, X any](fn func(A, B, C) X) func(B, A, C) X {
	return func(b B, a A, c C) X {
		return fn(a, b, c)
	}
}

func ArgRev1BCA[A, B, C, X any](fn func(A, B, C) X) func(B, C, A) X {
	return func(b B, c C, a A) X {
		return fn(a, b, c)
	}
}

func ArgRev1CAB[A, B, C, X any](fn func(A, B, C) X) func(C, A, B) X {
	return func(c C, a A, b B) X {
		return fn(a, b, c)
	}
}

func ArgRev1CBA[A, B, C, X any](fn func(A, B, C) X) func(C, B, A) X {
	return func(c C, b B, a A) X {
		return fn(a, b, c)
	}
}

func ArgRev2ACB[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(A, C, B) (X, Y) {
	return func(a A, c C, b B) (X, Y) {
		return fn(a, b, c)
	}
}

func ArgRev2BAC[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(B, A, C) (X, Y) {
	return func(b B, a A, c C) (X, Y) {
		return fn(a, b, c)
	}
}

func ArgRev2BCA[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(B, C, A) (X, Y) {
	return func(b B, c C, a A) (X, Y) {
		return fn(a, b, c)
	}
}

func ArgRev2CAB[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(C, A, B) (X, Y) {
	return func(c C, a A, b B) (X, Y) {
		return fn(a, b, c)
	}
}

func ArgRev2CBA[A, B, C, X, Y any](fn func(A, B, C) (X, Y)) func(C, B, A) (X, Y) {
	return func(c C, b B, a A) (X, Y) {
		return fn(a, b, c)
	}
}

func ArgRev3ACB[A, B, C, X, Y, Z any](fn func(A, B, C) (X, Y, Z)) func(A, C, B) (X, Y, Z) {
	return func(a A, c C, b B) (X, Y, Z) {
		return fn(a, b, c)
	}
}

func ArgRev3BAC[A, B, C, X, Y, Z any](fn func(A, B, C) (X, Y, Z)) func(B, A, C) (X, Y, Z) {
	return func(b B, a A, c C) (X, Y, Z) {
		return fn(a, b, c)
	}
}

func ArgRev3BCA[A, B, C, X, Y, Z any](fn func(A, B, C) (X, Y, Z)) func(B, C, A) (X, Y, Z) {
	return func(b B, c C, a A) (X, Y, Z) {
		return fn(a, b, c)
	}
}

func ArgRev3CAB[A, B, C, X, Y, Z any](fn func(A, B, C) (X, Y, Z)) func(C, A, B) (X, Y, Z) {
	return func(c C, a A, b B) (X, Y, Z) {
		return fn(a, b, c)
	}
}

func ArgRev3CBA[A, B, C, X, Y, Z any](fn func(A, B, C) (X, Y, Z)) func(C, B, A) (X, Y, Z) {
	return func(c C, b B, a A) (X, Y, Z) {
		return fn(a, b, c)
	}
}

func ArgRev0ABDC[A, B, C, D any](fn func(A, B, C, D)) func(A, B, D, C) {
	return func(a A, b B, d D, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0ACBD[A, B, C, D any](fn func(A, B, C, D)) func(A, C, B, D) {
	return func(a A, c C, b B, d D) {
		fn(a, b, c, d)
	}
}

func ArgRev0ACDB[A, B, C, D any](fn func(A, B, C, D)) func(A, C, D, B) {
	return func(a A, c C, d D, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0ADBC[A, B, C, D any](fn func(A, B, C, D)) func(A, D, B, C) {
	return func(a A, d D, b B, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0ADCB[A, B, C, D any](fn func(A, B, C, D)) func(A, D, C, B) {
	return func(a A, d D, c C, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0BACD[A, B, C, D any](fn func(A, B, C, D)) func(B, A, C, D) {
	return func(b B, a A, c C, d D) {
		fn(a, b, c, d)
	}
}

func ArgRev0BADC[A, B, C, D any](fn func(A, B, C, D)) func(B, A, D, C) {
	return func(b B, a A, d D, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0BCAD[A, B, C, D any](fn func(A, B, C, D)) func(B, C, A, D) {
	return func(b B, c C, a A, d D) {
		fn(a, b, c, d)
	}
}

func ArgRev0BCDA[A, B, C, D any](fn func(A, B, C, D)) func(B, C, D, A) {
	return func(b B, c C, d D, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev0BDAC[A, B, C, D any](fn func(A, B, C, D)) func(B, D, A, C) {
	return func(b B, d D, a A, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0BDCA[A, B, C, D any](fn func(A, B, C, D)) func(B, D, C, A) {
	return func(b B, d D, c C, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev0CABD[A, B, C, D any](fn func(A, B, C, D)) func(C, A, B, D) {
	return func(c C, a A, b B, d D) {
		fn(a, b, c, d)
	}
}

func ArgRev0CADB[A, B, C, D any](fn func(A, B, C, D)) func(C, A, D, B) {
	return func(c C, a A, d D, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0CBAD[A, B, C, D any](fn func(A, B, C, D)) func(C, B, A, D) {
	return func(c C, b B, a A, d D) {
		fn(a, b, c, d)
	}
}

func ArgRev0CBDA[A, B, C, D any](fn func(A, B, C, D)) func(C, B, D, A) {
	return func(c C, b B, d D, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev0CDAB[A, B, C, D any](fn func(A, B, C, D)) func(C, D, A, B) {
	return func(c C, d D, a A, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0CDBA[A, B, C, D any](fn func(A, B, C, D)) func(C, D, B, A) {
	return func(c C, d D, b B, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev0DABC[A, B, C, D any](fn func(A, B, C, D)) func(D, A, B, C) {
	return func(d D, a A, b B, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0DACB[A, B, C, D any](fn func(A, B, C, D)) func(D, A, C, B) {
	return func(d D, a A, c C, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0DBAC[A, B, C, D any](fn func(A, B, C, D)) func(D, B, A, C) {
	return func(d D, b B, a A, c C) {
		fn(a, b, c, d)
	}
}

func ArgRev0DBCA[A, B, C, D any](fn func(A, B, C, D)) func(D, B, C, A) {
	return func(d D, b B, c C, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev0DCAB[A, B, C, D any](fn func(A, B, C, D)) func(D, C, A, B) {
	return func(d D, c C, a A, b B) {
		fn(a, b, c, d)
	}
}

func ArgRev0DCBA[A, B, C, D any](fn func(A, B, C, D)) func(D, C, B, A) {
	return func(d D, c C, b B, a A) {
		fn(a, b, c, d)
	}
}

func ArgRev1ABDC[A, B, C, D, X any](fn func(A, B, C, D) X) func(A, B, D, C) X {
	return func(a A, b B, d D, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1ACBD[A, B, C, D, X any](fn func(A, B, C, D) X) func(A, C, B, D) X {
	return func(a A, c C, b B, d D) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1ACDB[A, B, C, D, X any](fn func(A, B, C, D) X) func(A, C, D, B) X {
	return func(a A, c C, d D, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1ADBC[A, B, C, D, X any](fn func(A, B, C, D) X) func(A, D, B, C) X {
	return func(a A, d D, b B, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1ADCB[A, B, C, D, X any](fn func(A, B, C, D) X) func(A, D, C, B) X {
	return func(a A, d D, c C, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BACD[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, A, C, D) X {
	return func(b B, a A, c C, d D) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BADC[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, A, D, C) X {
	return func(b B, a A, d D, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BCAD[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, C, A, D) X {
	return func(b B, c C, a A, d D) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BCDA[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, C, D, A) X {
	return func(b B, c C, d D, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BDAC[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, D, A, C) X {
	return func(b B, d D, a A, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1BDCA[A, B, C, D, X any](fn func(A, B, C, D) X) func(B, D, C, A) X {
	return func(b B, d D, c C, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CABD[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, A, B, D) X {
	return func(c C, a A, b B, d D) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CADB[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, A, D, B) X {
	return func(c C, a A, d D, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CBAD[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, B, A, D) X {
	return func(c C, b B, a A, d D) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CBDA[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, B, D, A) X {
	return func(c C, b B, d D, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CDAB[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, D, A, B) X {
	return func(c C, d D, a A, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1CDBA[A, B, C, D, X any](fn func(A, B, C, D) X) func(C, D, B, A) X {
	return func(c C, d D, b B, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DABC[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, A, B, C) X {
	return func(d D, a A, b B, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DACB[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, A, C, B) X {
	return func(d D, a A, c C, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DBAC[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, B, A, C) X {
	return func(d D, b B, a A, c C) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DBCA[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, B, C, A) X {
	return func(d D, b B, c C, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DCAB[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, C, A, B) X {
	return func(d D, c C, a A, b B) X {
		return fn(a, b, c, d)
	}
}

func ArgRev1DCBA[A, B, C, D, X any](fn func(A, B, C, D) X) func(D, C, B, A) X {
	return func(d D, c C, b B, a A) X {
		return fn(a, b, c, d)
	}
}

func ArgRev2ABDC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A, B, D, C) (X, Y) {
	return func(a A, b B, d D, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2ACBD[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A, C, B, D) (X, Y) {
	return func(a A, c C, b B, d D) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2ACDB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A, C, D, B) (X, Y) {
	return func(a A, c C, d D, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2ADBC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A, D, B, C) (X, Y) {
	return func(a A, d D, b B, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2ADCB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(A, D, C, B) (X, Y) {
	return func(a A, d D, c C, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BACD[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, A, C, D) (X, Y) {
	return func(b B, a A, c C, d D) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BADC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, A, D, C) (X, Y) {
	return func(b B, a A, d D, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BCAD[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, C, A, D) (X, Y) {
	return func(b B, c C, a A, d D) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BCDA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, C, D, A) (X, Y) {
	return func(b B, c C, d D, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BDAC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, D, A, C) (X, Y) {
	return func(b B, d D, a A, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2BDCA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(B, D, C, A) (X, Y) {
	return func(b B, d D, c C, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CABD[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, A, B, D) (X, Y) {
	return func(c C, a A, b B, d D) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CADB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, A, D, B) (X, Y) {
	return func(c C, a A, d D, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CBAD[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, B, A, D) (X, Y) {
	return func(c C, b B, a A, d D) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CBDA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, B, D, A) (X, Y) {
	return func(c C, b B, d D, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CDAB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, D, A, B) (X, Y) {
	return func(c C, d D, a A, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2CDBA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(C, D, B, A) (X, Y) {
	return func(c C, d D, b B, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DABC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, A, B, C) (X, Y) {
	return func(d D, a A, b B, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DACB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, A, C, B) (X, Y) {
	return func(d D, a A, c C, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DBAC[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, B, A, C) (X, Y) {
	return func(d D, b B, a A, c C) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DBCA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, B, C, A) (X, Y) {
	return func(d D, b B, c C, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DCAB[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, C, A, B) (X, Y) {
	return func(d D, c C, a A, b B) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev2DCBA[A, B, C, D, X, Y any](fn func(A, B, C, D) (X, Y)) func(D, C, B, A) (X, Y) {
	return func(d D, c C, b B, a A) (X, Y) {
		return fn(a, b, c, d)
	}
}

func ArgRev3ABDC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(A, B, D, C) (X, Y, Z) {
	return func(a A, b B, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3ACBD[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(A, C, B, D) (X, Y, Z) {
	return func(a A, c C, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3ACDB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(A, C, D, B) (X, Y, Z) {
	return func(a A, c C, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3ADBC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(A, D, B, C) (X, Y, Z) {
	return func(a A, d D, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3ADCB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(A, D, C, B) (X, Y, Z) {
	return func(a A, d D, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BACD[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, A, C, D) (X, Y, Z) {
	return func(b B, a A, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BADC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, A, D, C) (X, Y, Z) {
	return func(b B, a A, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BCAD[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, C, A, D) (X, Y, Z) {
	return func(b B, c C, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BCDA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, C, D, A) (X, Y, Z) {
	return func(b B, c C, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BDAC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, D, A, C) (X, Y, Z) {
	return func(b B, d D, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3BDCA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(B, D, C, A) (X, Y, Z) {
	return func(b B, d D, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CABD[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, A, B, D) (X, Y, Z) {
	return func(c C, a A, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CADB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, A, D, B) (X, Y, Z) {
	return func(c C, a A, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CBAD[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, B, A, D) (X, Y, Z) {
	return func(c C, b B, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CBDA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, B, D, A) (X, Y, Z) {
	return func(c C, b B, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CDAB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, D, A, B) (X, Y, Z) {
	return func(c C, d D, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3CDBA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(C, D, B, A) (X, Y, Z) {
	return func(c C, d D, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DABC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, A, B, C) (X, Y, Z) {
	return func(d D, a A, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DACB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, A, C, B) (X, Y, Z) {
	return func(d D, a A, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DBAC[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, B, A, C) (X, Y, Z) {
	return func(d D, b B, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DBCA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, B, C, A) (X, Y, Z) {
	return func(d D, b B, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DCAB[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, C, A, B) (X, Y, Z) {
	return func(d D, c C, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev3DCBA[A, B, C, D, X, Y, Z any](fn func(A, B, C, D) (X, Y, Z)) func(D, C, B, A) (X, Y, Z) {
	return func(d D, c C, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d)
	}
}

func ArgRev0ABCED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, B, C, E, D) {
	return func(a A, b B, c C, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ABDCE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, B, D, C, E) {
	return func(a A, b B, d D, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ABDEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, B, D, E, C) {
	return func(a A, b B, d D, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ABECD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, B, E, C, D) {
	return func(a A, b B, e E, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ABEDC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, B, E, D, C) {
	return func(a A, b B, e E, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACBDE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, B, D, E) {
	return func(a A, c C, b B, d D, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACBED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, B, E, D) {
	return func(a A, c C, b B, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACDBE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, D, B, E) {
	return func(a A, c C, d D, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACDEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, D, E, B) {
	return func(a A, c C, d D, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACEBD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, E, B, D) {
	return func(a A, c C, e E, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ACEDB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, C, E, D, B) {
	return func(a A, c C, e E, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADBCE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, B, C, E) {
	return func(a A, d D, b B, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADBEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, B, E, C) {
	return func(a A, d D, b B, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADCBE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, C, B, E) {
	return func(a A, d D, c C, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADCEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, C, E, B) {
	return func(a A, d D, c C, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADEBC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, E, B, C) {
	return func(a A, d D, e E, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ADECB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, D, E, C, B) {
	return func(a A, d D, e E, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AEBCD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, B, C, D) {
	return func(a A, e E, b B, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AEBDC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, B, D, C) {
	return func(a A, e E, b B, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AECBD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, C, B, D) {
	return func(a A, e E, c C, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AECDB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, C, D, B) {
	return func(a A, e E, c C, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AEDBC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, D, B, C) {
	return func(a A, e E, d D, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0AEDCB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(A, E, D, C, B) {
	return func(a A, e E, d D, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BACDE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, C, D, E) {
	return func(b B, a A, c C, d D, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BACED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, C, E, D) {
	return func(b B, a A, c C, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BADCE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, D, C, E) {
	return func(b B, a A, d D, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BADEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, D, E, C) {
	return func(b B, a A, d D, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BAECD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, E, C, D) {
	return func(b B, a A, e E, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BAEDC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, A, E, D, C) {
	return func(b B, a A, e E, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCADE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, A, D, E) {
	return func(b B, c C, a A, d D, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCAED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, A, E, D) {
	return func(b B, c C, a A, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCDAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, D, A, E) {
	return func(b B, c C, d D, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCDEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, D, E, A) {
	return func(b B, c C, d D, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCEAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, E, A, D) {
	return func(b B, c C, e E, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BCEDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, C, E, D, A) {
	return func(b B, c C, e E, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDACE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, A, C, E) {
	return func(b B, d D, a A, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDAEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, A, E, C) {
	return func(b B, d D, a A, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDCAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, C, A, E) {
	return func(b B, d D, c C, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDCEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, C, E, A) {
	return func(b B, d D, c C, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDEAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, E, A, C) {
	return func(b B, d D, e E, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BDECA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, D, E, C, A) {
	return func(b B, d D, e E, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BEACD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, A, C, D) {
	return func(b B, e E, a A, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BEADC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, A, D, C) {
	return func(b B, e E, a A, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BECAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, C, A, D) {
	return func(b B, e E, c C, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BECDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, C, D, A) {
	return func(b B, e E, c C, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BEDAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, D, A, C) {
	return func(b B, e E, d D, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0BEDCA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(B, E, D, C, A) {
	return func(b B, e E, d D, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CABDE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, B, D, E) {
	return func(c C, a A, b B, d D, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CABED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, B, E, D) {
	return func(c C, a A, b B, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CADBE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, D, B, E) {
	return func(c C, a A, d D, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CADEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, D, E, B) {
	return func(c C, a A, d D, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CAEBD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, E, B, D) {
	return func(c C, a A, e E, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CAEDB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, A, E, D, B) {
	return func(c C, a A, e E, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBADE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, A, D, E) {
	return func(c C, b B, a A, d D, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBAED[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, A, E, D) {
	return func(c C, b B, a A, e E, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBDAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, D, A, E) {
	return func(c C, b B, d D, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBDEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, D, E, A) {
	return func(c C, b B, d D, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBEAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, E, A, D) {
	return func(c C, b B, e E, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CBEDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, B, E, D, A) {
	return func(c C, b B, e E, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDABE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, A, B, E) {
	return func(c C, d D, a A, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDAEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, A, E, B) {
	return func(c C, d D, a A, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDBAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, B, A, E) {
	return func(c C, d D, b B, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDBEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, B, E, A) {
	return func(c C, d D, b B, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDEAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, E, A, B) {
	return func(c C, d D, e E, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CDEBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, D, E, B, A) {
	return func(c C, d D, e E, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEABD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, A, B, D) {
	return func(c C, e E, a A, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEADB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, A, D, B) {
	return func(c C, e E, a A, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEBAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, B, A, D) {
	return func(c C, e E, b B, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEBDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, B, D, A) {
	return func(c C, e E, b B, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEDAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, D, A, B) {
	return func(c C, e E, d D, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0CEDBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(C, E, D, B, A) {
	return func(c C, e E, d D, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DABCE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, B, C, E) {
	return func(d D, a A, b B, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DABEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, B, E, C) {
	return func(d D, a A, b B, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DACBE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, C, B, E) {
	return func(d D, a A, c C, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DACEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, C, E, B) {
	return func(d D, a A, c C, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DAEBC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, E, B, C) {
	return func(d D, a A, e E, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DAECB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, A, E, C, B) {
	return func(d D, a A, e E, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBACE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, A, C, E) {
	return func(d D, b B, a A, c C, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBAEC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, A, E, C) {
	return func(d D, b B, a A, e E, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBCAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, C, A, E) {
	return func(d D, b B, c C, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBCEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, C, E, A) {
	return func(d D, b B, c C, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBEAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, E, A, C) {
	return func(d D, b B, e E, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DBECA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, B, E, C, A) {
	return func(d D, b B, e E, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCABE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, A, B, E) {
	return func(d D, c C, a A, b B, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCAEB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, A, E, B) {
	return func(d D, c C, a A, e E, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCBAE[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, B, A, E) {
	return func(d D, c C, b B, a A, e E) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCBEA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, B, E, A) {
	return func(d D, c C, b B, e E, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCEAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, E, A, B) {
	return func(d D, c C, e E, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DCEBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, C, E, B, A) {
	return func(d D, c C, e E, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DEABC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, A, B, C) {
	return func(d D, e E, a A, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DEACB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, A, C, B) {
	return func(d D, e E, a A, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DEBAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, B, A, C) {
	return func(d D, e E, b B, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DEBCA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, B, C, A) {
	return func(d D, e E, b B, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DECAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, C, A, B) {
	return func(d D, e E, c C, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0DECBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(D, E, C, B, A) {
	return func(d D, e E, c C, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EABCD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, B, C, D) {
	return func(e E, a A, b B, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EABDC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, B, D, C) {
	return func(e E, a A, b B, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EACBD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, C, B, D) {
	return func(e E, a A, c C, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EACDB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, C, D, B) {
	return func(e E, a A, c C, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EADBC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, D, B, C) {
	return func(e E, a A, d D, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EADCB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, A, D, C, B) {
	return func(e E, a A, d D, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBACD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, A, C, D) {
	return func(e E, b B, a A, c C, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBADC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, A, D, C) {
	return func(e E, b B, a A, d D, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBCAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, C, A, D) {
	return func(e E, b B, c C, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBCDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, C, D, A) {
	return func(e E, b B, c C, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBDAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, D, A, C) {
	return func(e E, b B, d D, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EBDCA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, B, D, C, A) {
	return func(e E, b B, d D, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECABD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, A, B, D) {
	return func(e E, c C, a A, b B, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECADB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, A, D, B) {
	return func(e E, c C, a A, d D, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECBAD[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, B, A, D) {
	return func(e E, c C, b B, a A, d D) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECBDA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, B, D, A) {
	return func(e E, c C, b B, d D, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECDAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, D, A, B) {
	return func(e E, c C, d D, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0ECDBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, C, D, B, A) {
	return func(e E, c C, d D, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDABC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, A, B, C) {
	return func(e E, d D, a A, b B, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDACB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, A, C, B) {
	return func(e E, d D, a A, c C, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDBAC[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, B, A, C) {
	return func(e E, d D, b B, a A, c C) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDBCA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, B, C, A) {
	return func(e E, d D, b B, c C, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDCAB[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, C, A, B) {
	return func(e E, d D, c C, a A, b B) {
		fn(a, b, c, d, e)
	}
}

func ArgRev0EDCBA[A, B, C, D, E any](fn func(A, B, C, D, E)) func(E, D, C, B, A) {
	return func(e E, d D, c C, b B, a A) {
		fn(a, b, c, d, e)
	}
}

func ArgRev1ABCED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, B, C, E, D) X {
	return func(a A, b B, c C, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ABDCE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, B, D, C, E) X {
	return func(a A, b B, d D, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ABDEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, B, D, E, C) X {
	return func(a A, b B, d D, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ABECD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, B, E, C, D) X {
	return func(a A, b B, e E, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ABEDC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, B, E, D, C) X {
	return func(a A, b B, e E, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACBDE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, B, D, E) X {
	return func(a A, c C, b B, d D, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACBED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, B, E, D) X {
	return func(a A, c C, b B, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACDBE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, D, B, E) X {
	return func(a A, c C, d D, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACDEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, D, E, B) X {
	return func(a A, c C, d D, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACEBD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, E, B, D) X {
	return func(a A, c C, e E, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ACEDB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, C, E, D, B) X {
	return func(a A, c C, e E, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADBCE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, B, C, E) X {
	return func(a A, d D, b B, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADBEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, B, E, C) X {
	return func(a A, d D, b B, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADCBE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, C, B, E) X {
	return func(a A, d D, c C, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADCEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, C, E, B) X {
	return func(a A, d D, c C, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADEBC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, E, B, C) X {
	return func(a A, d D, e E, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ADECB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, D, E, C, B) X {
	return func(a A, d D, e E, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AEBCD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, B, C, D) X {
	return func(a A, e E, b B, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AEBDC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, B, D, C) X {
	return func(a A, e E, b B, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AECBD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, C, B, D) X {
	return func(a A, e E, c C, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AECDB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, C, D, B) X {
	return func(a A, e E, c C, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AEDBC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, D, B, C) X {
	return func(a A, e E, d D, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1AEDCB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(A, E, D, C, B) X {
	return func(a A, e E, d D, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BACDE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, C, D, E) X {
	return func(b B, a A, c C, d D, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BACED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, C, E, D) X {
	return func(b B, a A, c C, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BADCE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, D, C, E) X {
	return func(b B, a A, d D, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BADEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, D, E, C) X {
	return func(b B, a A, d D, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BAECD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, E, C, D) X {
	return func(b B, a A, e E, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BAEDC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, A, E, D, C) X {
	return func(b B, a A, e E, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCADE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, A, D, E) X {
	return func(b B, c C, a A, d D, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCAED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, A, E, D) X {
	return func(b B, c C, a A, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCDAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, D, A, E) X {
	return func(b B, c C, d D, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCDEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, D, E, A) X {
	return func(b B, c C, d D, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCEAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, E, A, D) X {
	return func(b B, c C, e E, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BCEDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, C, E, D, A) X {
	return func(b B, c C, e E, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDACE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, A, C, E) X {
	return func(b B, d D, a A, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDAEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, A, E, C) X {
	return func(b B, d D, a A, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDCAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, C, A, E) X {
	return func(b B, d D, c C, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDCEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, C, E, A) X {
	return func(b B, d D, c C, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDEAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, E, A, C) X {
	return func(b B, d D, e E, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BDECA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, D, E, C, A) X {
	return func(b B, d D, e E, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BEACD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, A, C, D) X {
	return func(b B, e E, a A, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BEADC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, A, D, C) X {
	return func(b B, e E, a A, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BECAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, C, A, D) X {
	return func(b B, e E, c C, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BECDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, C, D, A) X {
	return func(b B, e E, c C, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BEDAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, D, A, C) X {
	return func(b B, e E, d D, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1BEDCA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(B, E, D, C, A) X {
	return func(b B, e E, d D, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CABDE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, B, D, E) X {
	return func(c C, a A, b B, d D, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CABED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, B, E, D) X {
	return func(c C, a A, b B, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CADBE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, D, B, E) X {
	return func(c C, a A, d D, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CADEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, D, E, B) X {
	return func(c C, a A, d D, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CAEBD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, E, B, D) X {
	return func(c C, a A, e E, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CAEDB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, A, E, D, B) X {
	return func(c C, a A, e E, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBADE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, A, D, E) X {
	return func(c C, b B, a A, d D, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBAED[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, A, E, D) X {
	return func(c C, b B, a A, e E, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBDAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, D, A, E) X {
	return func(c C, b B, d D, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBDEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, D, E, A) X {
	return func(c C, b B, d D, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBEAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, E, A, D) X {
	return func(c C, b B, e E, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CBEDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, B, E, D, A) X {
	return func(c C, b B, e E, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDABE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, A, B, E) X {
	return func(c C, d D, a A, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDAEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, A, E, B) X {
	return func(c C, d D, a A, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDBAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, B, A, E) X {
	return func(c C, d D, b B, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDBEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, B, E, A) X {
	return func(c C, d D, b B, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDEAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, E, A, B) X {
	return func(c C, d D, e E, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CDEBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, D, E, B, A) X {
	return func(c C, d D, e E, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEABD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, A, B, D) X {
	return func(c C, e E, a A, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEADB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, A, D, B) X {
	return func(c C, e E, a A, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEBAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, B, A, D) X {
	return func(c C, e E, b B, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEBDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, B, D, A) X {
	return func(c C, e E, b B, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEDAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, D, A, B) X {
	return func(c C, e E, d D, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1CEDBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(C, E, D, B, A) X {
	return func(c C, e E, d D, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DABCE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, B, C, E) X {
	return func(d D, a A, b B, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DABEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, B, E, C) X {
	return func(d D, a A, b B, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DACBE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, C, B, E) X {
	return func(d D, a A, c C, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DACEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, C, E, B) X {
	return func(d D, a A, c C, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DAEBC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, E, B, C) X {
	return func(d D, a A, e E, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DAECB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, A, E, C, B) X {
	return func(d D, a A, e E, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBACE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, A, C, E) X {
	return func(d D, b B, a A, c C, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBAEC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, A, E, C) X {
	return func(d D, b B, a A, e E, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBCAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, C, A, E) X {
	return func(d D, b B, c C, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBCEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, C, E, A) X {
	return func(d D, b B, c C, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBEAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, E, A, C) X {
	return func(d D, b B, e E, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DBECA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, B, E, C, A) X {
	return func(d D, b B, e E, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCABE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, A, B, E) X {
	return func(d D, c C, a A, b B, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCAEB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, A, E, B) X {
	return func(d D, c C, a A, e E, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCBAE[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, B, A, E) X {
	return func(d D, c C, b B, a A, e E) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCBEA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, B, E, A) X {
	return func(d D, c C, b B, e E, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCEAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, E, A, B) X {
	return func(d D, c C, e E, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DCEBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, C, E, B, A) X {
	return func(d D, c C, e E, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DEABC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, A, B, C) X {
	return func(d D, e E, a A, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DEACB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, A, C, B) X {
	return func(d D, e E, a A, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DEBAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, B, A, C) X {
	return func(d D, e E, b B, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DEBCA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, B, C, A) X {
	return func(d D, e E, b B, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DECAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, C, A, B) X {
	return func(d D, e E, c C, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1DECBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(D, E, C, B, A) X {
	return func(d D, e E, c C, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EABCD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, B, C, D) X {
	return func(e E, a A, b B, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EABDC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, B, D, C) X {
	return func(e E, a A, b B, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EACBD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, C, B, D) X {
	return func(e E, a A, c C, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EACDB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, C, D, B) X {
	return func(e E, a A, c C, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EADBC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, D, B, C) X {
	return func(e E, a A, d D, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EADCB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, A, D, C, B) X {
	return func(e E, a A, d D, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBACD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, A, C, D) X {
	return func(e E, b B, a A, c C, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBADC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, A, D, C) X {
	return func(e E, b B, a A, d D, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBCAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, C, A, D) X {
	return func(e E, b B, c C, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBCDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, C, D, A) X {
	return func(e E, b B, c C, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBDAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, D, A, C) X {
	return func(e E, b B, d D, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EBDCA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, B, D, C, A) X {
	return func(e E, b B, d D, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECABD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, A, B, D) X {
	return func(e E, c C, a A, b B, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECADB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, A, D, B) X {
	return func(e E, c C, a A, d D, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECBAD[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, B, A, D) X {
	return func(e E, c C, b B, a A, d D) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECBDA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, B, D, A) X {
	return func(e E, c C, b B, d D, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECDAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, D, A, B) X {
	return func(e E, c C, d D, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1ECDBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, C, D, B, A) X {
	return func(e E, c C, d D, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDABC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, A, B, C) X {
	return func(e E, d D, a A, b B, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDACB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, A, C, B) X {
	return func(e E, d D, a A, c C, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDBAC[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, B, A, C) X {
	return func(e E, d D, b B, a A, c C) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDBCA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, B, C, A) X {
	return func(e E, d D, b B, c C, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDCAB[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, C, A, B) X {
	return func(e E, d D, c C, a A, b B) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev1EDCBA[A, B, C, D, E, X any](fn func(A, B, C, D, E) X) func(E, D, C, B, A) X {
	return func(e E, d D, c C, b B, a A) X {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ABCED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, B, C, E, D) (X, Y) {
	return func(a A, b B, c C, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ABDCE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, B, D, C, E) (X, Y) {
	return func(a A, b B, d D, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ABDEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, B, D, E, C) (X, Y) {
	return func(a A, b B, d D, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ABECD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, B, E, C, D) (X, Y) {
	return func(a A, b B, e E, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ABEDC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, B, E, D, C) (X, Y) {
	return func(a A, b B, e E, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACBDE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, B, D, E) (X, Y) {
	return func(a A, c C, b B, d D, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACBED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, B, E, D) (X, Y) {
	return func(a A, c C, b B, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACDBE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, D, B, E) (X, Y) {
	return func(a A, c C, d D, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACDEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, D, E, B) (X, Y) {
	return func(a A, c C, d D, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACEBD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, E, B, D) (X, Y) {
	return func(a A, c C, e E, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ACEDB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, C, E, D, B) (X, Y) {
	return func(a A, c C, e E, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADBCE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, B, C, E) (X, Y) {
	return func(a A, d D, b B, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADBEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, B, E, C) (X, Y) {
	return func(a A, d D, b B, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADCBE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, C, B, E) (X, Y) {
	return func(a A, d D, c C, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADCEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, C, E, B) (X, Y) {
	return func(a A, d D, c C, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADEBC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, E, B, C) (X, Y) {
	return func(a A, d D, e E, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ADECB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, D, E, C, B) (X, Y) {
	return func(a A, d D, e E, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AEBCD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, B, C, D) (X, Y) {
	return func(a A, e E, b B, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AEBDC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, B, D, C) (X, Y) {
	return func(a A, e E, b B, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AECBD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, C, B, D) (X, Y) {
	return func(a A, e E, c C, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AECDB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, C, D, B) (X, Y) {
	return func(a A, e E, c C, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AEDBC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, D, B, C) (X, Y) {
	return func(a A, e E, d D, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2AEDCB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(A, E, D, C, B) (X, Y) {
	return func(a A, e E, d D, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BACDE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, C, D, E) (X, Y) {
	return func(b B, a A, c C, d D, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BACED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, C, E, D) (X, Y) {
	return func(b B, a A, c C, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BADCE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, D, C, E) (X, Y) {
	return func(b B, a A, d D, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BADEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, D, E, C) (X, Y) {
	return func(b B, a A, d D, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BAECD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, E, C, D) (X, Y) {
	return func(b B, a A, e E, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BAEDC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, A, E, D, C) (X, Y) {
	return func(b B, a A, e E, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCADE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, A, D, E) (X, Y) {
	return func(b B, c C, a A, d D, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCAED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, A, E, D) (X, Y) {
	return func(b B, c C, a A, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCDAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, D, A, E) (X, Y) {
	return func(b B, c C, d D, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCDEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, D, E, A) (X, Y) {
	return func(b B, c C, d D, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCEAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, E, A, D) (X, Y) {
	return func(b B, c C, e E, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BCEDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, C, E, D, A) (X, Y) {
	return func(b B, c C, e E, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDACE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, A, C, E) (X, Y) {
	return func(b B, d D, a A, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDAEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, A, E, C) (X, Y) {
	return func(b B, d D, a A, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDCAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, C, A, E) (X, Y) {
	return func(b B, d D, c C, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDCEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, C, E, A) (X, Y) {
	return func(b B, d D, c C, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDEAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, E, A, C) (X, Y) {
	return func(b B, d D, e E, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BDECA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, D, E, C, A) (X, Y) {
	return func(b B, d D, e E, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BEACD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, A, C, D) (X, Y) {
	return func(b B, e E, a A, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BEADC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, A, D, C) (X, Y) {
	return func(b B, e E, a A, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BECAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, C, A, D) (X, Y) {
	return func(b B, e E, c C, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BECDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, C, D, A) (X, Y) {
	return func(b B, e E, c C, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BEDAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, D, A, C) (X, Y) {
	return func(b B, e E, d D, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2BEDCA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(B, E, D, C, A) (X, Y) {
	return func(b B, e E, d D, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CABDE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, B, D, E) (X, Y) {
	return func(c C, a A, b B, d D, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CABED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, B, E, D) (X, Y) {
	return func(c C, a A, b B, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CADBE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, D, B, E) (X, Y) {
	return func(c C, a A, d D, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CADEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, D, E, B) (X, Y) {
	return func(c C, a A, d D, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CAEBD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, E, B, D) (X, Y) {
	return func(c C, a A, e E, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CAEDB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, A, E, D, B) (X, Y) {
	return func(c C, a A, e E, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBADE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, A, D, E) (X, Y) {
	return func(c C, b B, a A, d D, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBAED[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, A, E, D) (X, Y) {
	return func(c C, b B, a A, e E, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBDAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, D, A, E) (X, Y) {
	return func(c C, b B, d D, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBDEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, D, E, A) (X, Y) {
	return func(c C, b B, d D, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBEAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, E, A, D) (X, Y) {
	return func(c C, b B, e E, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CBEDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, B, E, D, A) (X, Y) {
	return func(c C, b B, e E, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDABE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, A, B, E) (X, Y) {
	return func(c C, d D, a A, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDAEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, A, E, B) (X, Y) {
	return func(c C, d D, a A, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDBAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, B, A, E) (X, Y) {
	return func(c C, d D, b B, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDBEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, B, E, A) (X, Y) {
	return func(c C, d D, b B, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDEAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, E, A, B) (X, Y) {
	return func(c C, d D, e E, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CDEBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, D, E, B, A) (X, Y) {
	return func(c C, d D, e E, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEABD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, A, B, D) (X, Y) {
	return func(c C, e E, a A, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEADB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, A, D, B) (X, Y) {
	return func(c C, e E, a A, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEBAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, B, A, D) (X, Y) {
	return func(c C, e E, b B, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEBDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, B, D, A) (X, Y) {
	return func(c C, e E, b B, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEDAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, D, A, B) (X, Y) {
	return func(c C, e E, d D, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2CEDBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(C, E, D, B, A) (X, Y) {
	return func(c C, e E, d D, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DABCE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, B, C, E) (X, Y) {
	return func(d D, a A, b B, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DABEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, B, E, C) (X, Y) {
	return func(d D, a A, b B, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DACBE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, C, B, E) (X, Y) {
	return func(d D, a A, c C, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DACEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, C, E, B) (X, Y) {
	return func(d D, a A, c C, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DAEBC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, E, B, C) (X, Y) {
	return func(d D, a A, e E, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DAECB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, A, E, C, B) (X, Y) {
	return func(d D, a A, e E, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBACE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, A, C, E) (X, Y) {
	return func(d D, b B, a A, c C, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBAEC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, A, E, C) (X, Y) {
	return func(d D, b B, a A, e E, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBCAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, C, A, E) (X, Y) {
	return func(d D, b B, c C, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBCEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, C, E, A) (X, Y) {
	return func(d D, b B, c C, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBEAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, E, A, C) (X, Y) {
	return func(d D, b B, e E, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DBECA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, B, E, C, A) (X, Y) {
	return func(d D, b B, e E, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCABE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, A, B, E) (X, Y) {
	return func(d D, c C, a A, b B, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCAEB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, A, E, B) (X, Y) {
	return func(d D, c C, a A, e E, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCBAE[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, B, A, E) (X, Y) {
	return func(d D, c C, b B, a A, e E) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCBEA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, B, E, A) (X, Y) {
	return func(d D, c C, b B, e E, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCEAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, E, A, B) (X, Y) {
	return func(d D, c C, e E, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DCEBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, C, E, B, A) (X, Y) {
	return func(d D, c C, e E, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DEABC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, A, B, C) (X, Y) {
	return func(d D, e E, a A, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DEACB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, A, C, B) (X, Y) {
	return func(d D, e E, a A, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DEBAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, B, A, C) (X, Y) {
	return func(d D, e E, b B, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DEBCA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, B, C, A) (X, Y) {
	return func(d D, e E, b B, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DECAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, C, A, B) (X, Y) {
	return func(d D, e E, c C, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2DECBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(D, E, C, B, A) (X, Y) {
	return func(d D, e E, c C, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EABCD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, B, C, D) (X, Y) {
	return func(e E, a A, b B, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EABDC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, B, D, C) (X, Y) {
	return func(e E, a A, b B, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EACBD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, C, B, D) (X, Y) {
	return func(e E, a A, c C, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EACDB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, C, D, B) (X, Y) {
	return func(e E, a A, c C, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EADBC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, D, B, C) (X, Y) {
	return func(e E, a A, d D, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EADCB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, A, D, C, B) (X, Y) {
	return func(e E, a A, d D, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBACD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, A, C, D) (X, Y) {
	return func(e E, b B, a A, c C, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBADC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, A, D, C) (X, Y) {
	return func(e E, b B, a A, d D, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBCAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, C, A, D) (X, Y) {
	return func(e E, b B, c C, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBCDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, C, D, A) (X, Y) {
	return func(e E, b B, c C, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBDAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, D, A, C) (X, Y) {
	return func(e E, b B, d D, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EBDCA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, B, D, C, A) (X, Y) {
	return func(e E, b B, d D, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECABD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, A, B, D) (X, Y) {
	return func(e E, c C, a A, b B, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECADB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, A, D, B) (X, Y) {
	return func(e E, c C, a A, d D, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECBAD[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, B, A, D) (X, Y) {
	return func(e E, c C, b B, a A, d D) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECBDA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, B, D, A) (X, Y) {
	return func(e E, c C, b B, d D, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECDAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, D, A, B) (X, Y) {
	return func(e E, c C, d D, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2ECDBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, C, D, B, A) (X, Y) {
	return func(e E, c C, d D, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDABC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, A, B, C) (X, Y) {
	return func(e E, d D, a A, b B, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDACB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, A, C, B) (X, Y) {
	return func(e E, d D, a A, c C, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDBAC[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, B, A, C) (X, Y) {
	return func(e E, d D, b B, a A, c C) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDBCA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, B, C, A) (X, Y) {
	return func(e E, d D, b B, c C, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDCAB[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, C, A, B) (X, Y) {
	return func(e E, d D, c C, a A, b B) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev2EDCBA[A, B, C, D, E, X, Y any](fn func(A, B, C, D, E) (X, Y)) func(E, D, C, B, A) (X, Y) {
	return func(e E, d D, c C, b B, a A) (X, Y) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ABCED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, B, C, E, D) (X, Y, Z) {
	return func(a A, b B, c C, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ABDCE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, B, D, C, E) (X, Y, Z) {
	return func(a A, b B, d D, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ABDEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, B, D, E, C) (X, Y, Z) {
	return func(a A, b B, d D, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ABECD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, B, E, C, D) (X, Y, Z) {
	return func(a A, b B, e E, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ABEDC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, B, E, D, C) (X, Y, Z) {
	return func(a A, b B, e E, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACBDE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, B, D, E) (X, Y, Z) {
	return func(a A, c C, b B, d D, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACBED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, B, E, D) (X, Y, Z) {
	return func(a A, c C, b B, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACDBE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, D, B, E) (X, Y, Z) {
	return func(a A, c C, d D, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACDEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, D, E, B) (X, Y, Z) {
	return func(a A, c C, d D, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACEBD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, E, B, D) (X, Y, Z) {
	return func(a A, c C, e E, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ACEDB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, C, E, D, B) (X, Y, Z) {
	return func(a A, c C, e E, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADBCE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, B, C, E) (X, Y, Z) {
	return func(a A, d D, b B, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADBEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, B, E, C) (X, Y, Z) {
	return func(a A, d D, b B, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADCBE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, C, B, E) (X, Y, Z) {
	return func(a A, d D, c C, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADCEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, C, E, B) (X, Y, Z) {
	return func(a A, d D, c C, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADEBC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, E, B, C) (X, Y, Z) {
	return func(a A, d D, e E, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ADECB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, D, E, C, B) (X, Y, Z) {
	return func(a A, d D, e E, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AEBCD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, B, C, D) (X, Y, Z) {
	return func(a A, e E, b B, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AEBDC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, B, D, C) (X, Y, Z) {
	return func(a A, e E, b B, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AECBD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, C, B, D) (X, Y, Z) {
	return func(a A, e E, c C, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AECDB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, C, D, B) (X, Y, Z) {
	return func(a A, e E, c C, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AEDBC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, D, B, C) (X, Y, Z) {
	return func(a A, e E, d D, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3AEDCB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(A, E, D, C, B) (X, Y, Z) {
	return func(a A, e E, d D, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BACDE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, C, D, E) (X, Y, Z) {
	return func(b B, a A, c C, d D, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BACED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, C, E, D) (X, Y, Z) {
	return func(b B, a A, c C, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BADCE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, D, C, E) (X, Y, Z) {
	return func(b B, a A, d D, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BADEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, D, E, C) (X, Y, Z) {
	return func(b B, a A, d D, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BAECD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, E, C, D) (X, Y, Z) {
	return func(b B, a A, e E, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BAEDC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, A, E, D, C) (X, Y, Z) {
	return func(b B, a A, e E, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCADE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, A, D, E) (X, Y, Z) {
	return func(b B, c C, a A, d D, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCAED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, A, E, D) (X, Y, Z) {
	return func(b B, c C, a A, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCDAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, D, A, E) (X, Y, Z) {
	return func(b B, c C, d D, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCDEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, D, E, A) (X, Y, Z) {
	return func(b B, c C, d D, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCEAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, E, A, D) (X, Y, Z) {
	return func(b B, c C, e E, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BCEDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, C, E, D, A) (X, Y, Z) {
	return func(b B, c C, e E, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDACE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, A, C, E) (X, Y, Z) {
	return func(b B, d D, a A, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDAEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, A, E, C) (X, Y, Z) {
	return func(b B, d D, a A, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDCAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, C, A, E) (X, Y, Z) {
	return func(b B, d D, c C, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDCEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, C, E, A) (X, Y, Z) {
	return func(b B, d D, c C, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDEAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, E, A, C) (X, Y, Z) {
	return func(b B, d D, e E, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BDECA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, D, E, C, A) (X, Y, Z) {
	return func(b B, d D, e E, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BEACD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, A, C, D) (X, Y, Z) {
	return func(b B, e E, a A, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BEADC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, A, D, C) (X, Y, Z) {
	return func(b B, e E, a A, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BECAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, C, A, D) (X, Y, Z) {
	return func(b B, e E, c C, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BECDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, C, D, A) (X, Y, Z) {
	return func(b B, e E, c C, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BEDAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, D, A, C) (X, Y, Z) {
	return func(b B, e E, d D, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3BEDCA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(B, E, D, C, A) (X, Y, Z) {
	return func(b B, e E, d D, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CABDE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, B, D, E) (X, Y, Z) {
	return func(c C, a A, b B, d D, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CABED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, B, E, D) (X, Y, Z) {
	return func(c C, a A, b B, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CADBE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, D, B, E) (X, Y, Z) {
	return func(c C, a A, d D, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CADEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, D, E, B) (X, Y, Z) {
	return func(c C, a A, d D, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CAEBD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, E, B, D) (X, Y, Z) {
	return func(c C, a A, e E, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CAEDB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, A, E, D, B) (X, Y, Z) {
	return func(c C, a A, e E, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBADE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, A, D, E) (X, Y, Z) {
	return func(c C, b B, a A, d D, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBAED[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, A, E, D) (X, Y, Z) {
	return func(c C, b B, a A, e E, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBDAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, D, A, E) (X, Y, Z) {
	return func(c C, b B, d D, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBDEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, D, E, A) (X, Y, Z) {
	return func(c C, b B, d D, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBEAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, E, A, D) (X, Y, Z) {
	return func(c C, b B, e E, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CBEDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, B, E, D, A) (X, Y, Z) {
	return func(c C, b B, e E, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDABE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, A, B, E) (X, Y, Z) {
	return func(c C, d D, a A, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDAEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, A, E, B) (X, Y, Z) {
	return func(c C, d D, a A, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDBAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, B, A, E) (X, Y, Z) {
	return func(c C, d D, b B, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDBEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, B, E, A) (X, Y, Z) {
	return func(c C, d D, b B, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDEAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, E, A, B) (X, Y, Z) {
	return func(c C, d D, e E, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CDEBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, D, E, B, A) (X, Y, Z) {
	return func(c C, d D, e E, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEABD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, A, B, D) (X, Y, Z) {
	return func(c C, e E, a A, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEADB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, A, D, B) (X, Y, Z) {
	return func(c C, e E, a A, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEBAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, B, A, D) (X, Y, Z) {
	return func(c C, e E, b B, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEBDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, B, D, A) (X, Y, Z) {
	return func(c C, e E, b B, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEDAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, D, A, B) (X, Y, Z) {
	return func(c C, e E, d D, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3CEDBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(C, E, D, B, A) (X, Y, Z) {
	return func(c C, e E, d D, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DABCE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, B, C, E) (X, Y, Z) {
	return func(d D, a A, b B, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DABEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, B, E, C) (X, Y, Z) {
	return func(d D, a A, b B, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DACBE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, C, B, E) (X, Y, Z) {
	return func(d D, a A, c C, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DACEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, C, E, B) (X, Y, Z) {
	return func(d D, a A, c C, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DAEBC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, E, B, C) (X, Y, Z) {
	return func(d D, a A, e E, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DAECB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, A, E, C, B) (X, Y, Z) {
	return func(d D, a A, e E, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBACE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, A, C, E) (X, Y, Z) {
	return func(d D, b B, a A, c C, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBAEC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, A, E, C) (X, Y, Z) {
	return func(d D, b B, a A, e E, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBCAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, C, A, E) (X, Y, Z) {
	return func(d D, b B, c C, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBCEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, C, E, A) (X, Y, Z) {
	return func(d D, b B, c C, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBEAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, E, A, C) (X, Y, Z) {
	return func(d D, b B, e E, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DBECA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, B, E, C, A) (X, Y, Z) {
	return func(d D, b B, e E, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCABE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, A, B, E) (X, Y, Z) {
	return func(d D, c C, a A, b B, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCAEB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, A, E, B) (X, Y, Z) {
	return func(d D, c C, a A, e E, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCBAE[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, B, A, E) (X, Y, Z) {
	return func(d D, c C, b B, a A, e E) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCBEA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, B, E, A) (X, Y, Z) {
	return func(d D, c C, b B, e E, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCEAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, E, A, B) (X, Y, Z) {
	return func(d D, c C, e E, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DCEBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, C, E, B, A) (X, Y, Z) {
	return func(d D, c C, e E, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DEABC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, A, B, C) (X, Y, Z) {
	return func(d D, e E, a A, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DEACB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, A, C, B) (X, Y, Z) {
	return func(d D, e E, a A, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DEBAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, B, A, C) (X, Y, Z) {
	return func(d D, e E, b B, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DEBCA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, B, C, A) (X, Y, Z) {
	return func(d D, e E, b B, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DECAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, C, A, B) (X, Y, Z) {
	return func(d D, e E, c C, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3DECBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(D, E, C, B, A) (X, Y, Z) {
	return func(d D, e E, c C, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EABCD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, B, C, D) (X, Y, Z) {
	return func(e E, a A, b B, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EABDC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, B, D, C) (X, Y, Z) {
	return func(e E, a A, b B, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EACBD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, C, B, D) (X, Y, Z) {
	return func(e E, a A, c C, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EACDB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, C, D, B) (X, Y, Z) {
	return func(e E, a A, c C, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EADBC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, D, B, C) (X, Y, Z) {
	return func(e E, a A, d D, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EADCB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, A, D, C, B) (X, Y, Z) {
	return func(e E, a A, d D, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBACD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, A, C, D) (X, Y, Z) {
	return func(e E, b B, a A, c C, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBADC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, A, D, C) (X, Y, Z) {
	return func(e E, b B, a A, d D, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBCAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, C, A, D) (X, Y, Z) {
	return func(e E, b B, c C, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBCDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, C, D, A) (X, Y, Z) {
	return func(e E, b B, c C, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBDAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, D, A, C) (X, Y, Z) {
	return func(e E, b B, d D, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EBDCA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, B, D, C, A) (X, Y, Z) {
	return func(e E, b B, d D, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECABD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, A, B, D) (X, Y, Z) {
	return func(e E, c C, a A, b B, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECADB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, A, D, B) (X, Y, Z) {
	return func(e E, c C, a A, d D, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECBAD[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, B, A, D) (X, Y, Z) {
	return func(e E, c C, b B, a A, d D) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECBDA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, B, D, A) (X, Y, Z) {
	return func(e E, c C, b B, d D, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECDAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, D, A, B) (X, Y, Z) {
	return func(e E, c C, d D, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3ECDBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, C, D, B, A) (X, Y, Z) {
	return func(e E, c C, d D, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDABC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, A, B, C) (X, Y, Z) {
	return func(e E, d D, a A, b B, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDACB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, A, C, B) (X, Y, Z) {
	return func(e E, d D, a A, c C, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDBAC[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, B, A, C) (X, Y, Z) {
	return func(e E, d D, b B, a A, c C) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDBCA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, B, C, A) (X, Y, Z) {
	return func(e E, d D, b B, c C, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDCAB[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, C, A, B) (X, Y, Z) {
	return func(e E, d D, c C, a A, b B) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}

func ArgRev3EDCBA[A, B, C, D, E, X, Y, Z any](fn func(A, B, C, D, E) (X, Y, Z)) func(E, D, C, B, A) (X, Y, Z) {
	return func(e E, d D, c C, b B, a A) (X, Y, Z) {
		return fn(a, b, c, d, e)
	}
}
