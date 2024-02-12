package main

import (
	"fmt"
	"math"
)

func main() {

	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}

// Closure
func powinator(a float64) func() float64 {
	c := 0.0
	return func() float64 {
		c++                   // This is the power (stepenot na inkrementi - sekoja iteracija mu e +1 to est vneseniot broj vo samata funkcija na samiot stepen)
		return math.Pow(a, c) // Go zima brojot sto e vnesen vo funkcijata, i go zgolemuva za stepenot to est 2 na prva, koa ke se povika pak, 2 na vtora itn itn itn...
	}
}

func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}

func returnio() func() int {
	return func() int {
		return 42
	}
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, n int) string {
	sum := f(n)
	return fmt.Sprintf("%v is the area of a square", sum)
}
