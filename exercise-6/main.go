package main

import (
	"fmt"
	"math/rand"
)
func init() {
	fmt.Println("I run first")
}

func main() {
	x := rand.Intn(251)
	fmt.Println(x)
	if x <= 100 {
		fmt.Printf("if/else: Below 100 - x: %v\n", x)
	} else if x > 100 && x <= 200 {
		fmt.Printf("if/else: Between 101 & 200 - x: %v\n", x)
	} else {
		fmt.Printf("if/else: Above 200 - x: %v\n", x)
	}

	y := rand.Intn(2)
	if y == 0 {
		fmt.Printf("y is %v", y)
	}

	x = rand.Intn(251)

	switch {
	case x == 0:
		fmt.Printf("switch: x is 0 - x: %v\n", x)
	case x > 0 && x <= 100:
		fmt.Printf("switch: Below 100 - x: %v\n", x)
	case x > 100 && x <= 200:
		fmt.Printf("switch: Between 101 & 200 - x: %v\n", x)
	default:
		fmt.Printf("switch: Above 200 - x: %v\n", x)
	}

}
