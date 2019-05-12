package math

func intPow(a, b int) int {
	if b == 0 {
		return 1
	} else if b % 2 == 1 {
		return a * intPow(a, b - 1)
	} else {
		p := intPow(a, b/2)
		return p*p
	}
}

// IntPow calculates the power of base^exponent
func IntPow(base, exponent int) int {
	return intPow(base, exponent)
}
