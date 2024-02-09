package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("---------- iteration %v ----------\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer %v \t\t Inner: %v\n", i, j)
		}
	}
}
