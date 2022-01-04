package words

func Z(s string) []int {
	ss := []rune(s)
	z := make([]int, len(ss))

	z[0] = len(z)
	L, R := 0, 1

	for k := 1; k < len(ss); k++ {
		if k < R && z[k-L] < (R-k) {
			z[k] = z[k-L]
		} else {
			L = k
			if R < L {
				R = L
			}
			for R < len(ss) && ss[R] == ss[R-k] {
				R++
			}
			z[k] = R - L
		}
	}

	return z
}
