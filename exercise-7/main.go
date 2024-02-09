package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)
	fmt.Println(x, y)
	if x < 4 && y < 4 {
		fmt.Printf("x and y are less than 4 - x: %v y:\n%v", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("x and y are greater than 6 - x: %v y:\n%v", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is between 4 & 6 inclusive - x: %v", x)
	} else if y != 5 {
		fmt.Printf("y is not 5 - y:\n %v", y)
	} else {
		fmt.Printf("Nothing from the conditions happened - x: %v y:\n %v", x, y)
	}

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x and y are less than 4 - x: %v y:\n%v", x, y)
	case x > 6 && y > 6:
		fmt.Printf("x and y are greater than 6 - x: %v y:\n%v", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("x is between 4 & 6 inclusive - x: %v", x)
	case y != 5:
		fmt.Printf("y is not 5 - y:\n %v", y)
	default:
		fmt.Printf("Nothing from the conditions happened - x: %v y:\n %v", x, y)
	}
}
