package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v \t first for loop\n", i)
	}

	var y int
	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("i: %v\n", i)
		if i%2 != 0 {
			continue
		}
		i++
	}
	// Nested
	for i := 0; i < 5; i++ {
		fmt.Println("------")
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop: %v \t inner loop %v\n", i, j)
		}
	}
	// For range - slice
	xi := []int{42, 43, 44, 45, 46, 47, 48}
	for i, v := range xi {
		fmt.Println("Ranging over a slice", i, v)
	}
	// for range - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("Key: %s \t value: %v\n", k, v)
	}

}
