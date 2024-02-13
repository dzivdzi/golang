package main

import "fmt"

// Value semantics - this creates a COPY of the value but doesn't modify the value in the underlying address meaning once it is finished, the address still holds the old value before the modification
func addOne(v int) int {
	return v + 1
}

// Pointer semantics - this modifies the value in the address
func addOneP(v *int){
	*v += 1
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	y := &x

	fmt.Println(*y)

	// Value semantics
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	// Pointer semantics
	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

}
