package utils

func a(k int, x1, x2, x3, x4, x5 func() int) int {
	var b func() int
	b = func() int {
		k--
		return a(k, b, x1, x2, x3, x4)
	}
	if k <= 0 {
		return x4() + x5()
	}
	return b()
}

func ManOrBoy() int {
	x := func(i int) func() int { return func() int { return i } }
	return a(10, x(1), x(-1), x(-1), x(1), x(0))
}
