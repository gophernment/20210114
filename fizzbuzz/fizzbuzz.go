package fizzbuzz

import "fmt"

func Say(n int) string {
	if n == 3 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", n)
}
