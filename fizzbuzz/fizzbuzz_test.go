package fizzbuzz_test

import (
	"testing"

	"github.com/pallat/20210114/fizzbuzz"
)

func TestFizzBuzzGiven1(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q bug get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven2(t *testing.T) {
	given := 2
	want := "2"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q bug get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven3(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q bug get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven4(t *testing.T) {
	given := 4
	want := "4"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q bug get %q\n", given, want, get)
	}
}
func TestFizzBuzzGiven5(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.Say(given)

	if want != get {
		t.Errorf("given %d want %q bug get %q\n", given, want, get)
	}
}
