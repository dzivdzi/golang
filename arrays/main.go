package main

import "fmt"

func main() {
	// declare the array of type int - takes 7 elements
	var ni [7]int
	// assign a value of 42 to index 0
	ni[0] = 42

	// array literall - the short definition of creating an array
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	// if you don't know how many elems the array will have u use the ... when declaring
	b := [...]string{"and this", "is", "with", "an unknown", "number of elems"}
	fmt.Println(b)

}
