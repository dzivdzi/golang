package main

import "fmt"

var x = 40

const y = 41

func main() {
	z := 11
	fmt.Printf("The value of x is %v and the type is %T\n", x, x)
	fmt.Printf("The value of y is %v and the type is %T\n", y, y)
	fmt.Printf("The value of z is %v and the type is %T\n", z, z)
}
