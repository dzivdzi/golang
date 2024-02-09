package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for k, v := range xi {
		fmt.Printf("Index: %v \t\t Value: %v\n", k, v)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	age := m["James"]
	fmt.Println(age)
	ageQ := m["Q"]
	fmt.Println(ageQ)

	if _, ok := m["Q"]; !ok {
		fmt.Println(ok)
		fmt.Println("Can't find Q in the map")
	}
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")
		}
	}
}
