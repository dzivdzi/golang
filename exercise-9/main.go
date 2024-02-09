package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("--------- Iteration %v ---------\n", i)
		switch {
		case x == 0:
			fmt.Printf("X: %v \t\t Type: %T\n", x, x)
		case x == 1:
			fmt.Printf("X: %v \t\t Type: %T\n", x, x)
		case x == 2:
			fmt.Printf("X: %v \t\t Type: %T\n", x, x)
		case x == 3:
			fmt.Printf("X: %v \t\t Type: %T\n", x, x)
		case x == 4:
			fmt.Printf("X: %v \t\t Type: %T\n", x, x)
		}
	}

}
