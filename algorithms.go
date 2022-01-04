package words

func Z[T comparable](s []T) []int {
	z := make([]int, len(s))

	z[0] = len(z)
	L, R := 0, 1

	for k := 1; k < len(s); k++ {
		if k < R && z[k-L] < (R-k) {
			z[k] = z[k-L]
		} else {
			L = k
			if R < L {
				R = L
			}
			for R < len(s) && s[R] == s[R-k] {
				R++
			}
			z[k] = R - L
		}
	}

	return z
}
