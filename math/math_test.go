package math_test

import (
	"testing"

	"github.com/pallat/20210114/math"
)

func TestPower(t *testing.T) {
	givenB := 2
	givenX := 3

	want := 8
	get := math.Power(givenB, givenX)
	if want != get {
		t.Errorf("given b = %d and x =%d want %d but get %d\n", givenB, givenX, want, get)
	}
}
