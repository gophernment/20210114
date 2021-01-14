package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pallat/20210114/fizzbuzz"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/fizzbuzz/:number", fizzbuzzHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func fizzbuzzHandler(c echo.Context) error {
	number := c.Param("number")

	n, err := strconv.Atoi(number)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fizzbuzz.Say(n))
}

func myPrintln(n int) {
	time.Sleep(time.Millisecond * 100)

	fmt.Println(n)
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
