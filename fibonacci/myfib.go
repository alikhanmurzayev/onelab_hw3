package fibonacci

// Fibonacci returns fibonacci sequence
func Fibonacci() func() int {
	var n = 0
	var a, b = 0, 1
	return func() int {
		if n == 0 {
			n++
			return 0
		} else if n == 1{
			n++
			return b
		}


		c := a + b
		a = b
		b = c

		n++
		return b
	}
}