package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("------- Iteration %v -------\n", i)
		fmt.Println(i)
		x := rand.Intn(11)
		y := rand.Intn(11)
		if x < 4 && y < 4 {
			fmt.Printf("x and y are less than 4 - x: %v y: %v\n", x, y)
		} else if x > 6 && y > 6 {
			fmt.Printf("x and y are greater than 6 - x: %v y:%v\n", x, y)
		} else if x >= 4 && x <= 6 {
			fmt.Printf("x is between 4 & 6 inclusive - x: %v\n", x)
		} else if y != 5 {
			fmt.Printf("y is not 5 - y: %v\n", y)
		} else {
			fmt.Printf("Nothing from the conditions happened - x: %v y: %v\n", x, y)
		}
	}
}
