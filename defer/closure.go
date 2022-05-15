package defer_test

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a, a)
		a = a + 1
		return a
	}
}
