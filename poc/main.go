package main

import (
	"fmt"
	"log"
	"strconv"
)

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Triangle struct {
	Width  float64
	Length float64
}

func (r Triangle) Area() float64 {
	return r.Width * r.Length * 0.5
}

type Object interface {
	Area() float64
}

func printArea(r Object) {
	fmt.Println(r.Area())
}

type xxx bool

func (xxx) String() string {
	return "mama"
}
func main() {
	var x xxx
	fmt.Println(x)
}

func catchMe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	log.Fatal("boom")

	s := []int{}

	fmt.Println(s[1])
}

func xfunc(n int) {
	defer func() { fmt.Println(n) }()
	defer fmt.Println(n)
	n *= n
	fmt.Println(n)
}

func ExampleMap() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
		"d": "durian",
		"e": "elderberry",
		"f": "fig",
		"g": "guava",
	}

	delete(m, "d")

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func ExampleSlice() {
	var primes = [...]int{2, 3, 5, 7}
	var slice []int

	// if slice == nil {
	// 	fmt.Println("it's nil")
	// }

	slice = primes[:2]

	// slice = append(slice, 19, 23, 27)

	variadic(slice...)
}

func variadic(nums ...int) {
	fmt.Printf("%#v\n", nums)
}
