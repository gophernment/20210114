package math

func Power(b, x int) int {
	result := 1

	for i := 1; i <= x; i++ {
		result *= b
	}
	return result
}
