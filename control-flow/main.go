package main

import (
	"fmt"
	"math/rand"
)

// Init func is the only thing that runs before main

func init() {

	fmt.Println("I run before main!")
}

func main() {
	fmt.Println("First run")
	fmt.Println("Second run")

	x := 40
	y := 5
	fmt.Printf("x=%v\n y=%v\n", x, y)

	if x < 42 {
		fmt.Println("Less than 42")
	}

	if x < 42 {
		fmt.Println("Less than 42")
	} else {
		fmt.Println("Greater or equal than 42")
	}

	if x < 42 {
		fmt.Println("Less than 42")
	} else if x == 42 {
		fmt.Println("Yay it's 42")
	} else {
		fmt.Println("Greater than 42")
	}

	// statement;statement & comma, ok idioms
	// In the below example, the scope of Z is ONLY in the if/else block
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v, x is %v\n", z, x)
	}

	switch x {
	case 40:
		fmt.Printf("X: %v", x)
	case 41:
		fmt.Printf("X: %v", x)
	case 42:
		fmt.Printf("X: %v", x)
	case 43:
		fmt.Printf("X: %v", x)
	default:
		fmt.Printf("X: %v", x)
	}

	// fallthrough is a special reserved keyword - basically it states continue with the evaluation
	switch x {
	case 40:
		fmt.Printf("X: %v", x)
		fallthrough
	case 41:
		fmt.Printf("X: %v", x)
	case 42:
		fmt.Printf("X: %v", x)
	case 43:
		fmt.Printf("X: %v", x)
	default:
		fmt.Printf("X: %v", x)
	}

}
