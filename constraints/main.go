package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// If there is a ~ in front of the type, it will take every value of subtype int meaning that it will also take
// the below created "myAlias" type created because the underlying type of my alias is "int"
type myNumbers interface {
	~int | ~float64
}

// Square brackets after the func name specifies which types the function can take(Type Constraint)
// Above we are creating an interface to combine the types to shorten the code and write it cleaner(Type Set)
func addT[T myNumbers](a, b T) T {
	return a + b
}

// Type Alias (Name your types)
type myAlias int

func main() {
	fmt.Println(addT(2, 2))
	var n myAlias = 42
	fmt.Println(addT(n, 2))
}
