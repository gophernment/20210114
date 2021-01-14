package fizzbuzz

import "strconv"

func Say(n int) string {
	if n == 5 {
		return "Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(n)
}
