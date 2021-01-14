package main

import "github.com/pallat/20210114/math"

func main() {
	println(math.Power(2, 4))
}

func prime(n int) {
	for i := 1; i <= n; i++ {
		var count int
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			println(i)
		}
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func squareArea(a int) int {
	return a * a
}
