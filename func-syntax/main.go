package main

import "fmt"

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

type person struct {
	first string
}

func main() {
	foo()
	bar("Steve")
	s := aloha("Steve")
	fmt.Println(s)
	s1, n := dogYears("Steve", 40)
	fmt.Println(s1, n)

	// Variadic Parameters
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum is", x)

	// Unfirling a slice (decoupling, extracting all of the values from the slice)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := sum(xi...) // we use the 3 dots notation to tell GO to extract all of the integers from the slice of ints
	fmt.Println(z)

	// Defer (delay the execution until the outer surrounding function ends)
	// "Defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because:
	// The surrounding function executed a return statement OR
	// Reached the end of it's function body OR
	// because the corresponding gorouting is PANICKING
	defer foo()
	bar("executed")

	// Method
	m := person{
		first: "James",
	}

	m.speak()

}

// Methods
func (p person) speak() {
	fmt.Println("I am", p.first)
}

// no params no returns
func foo() {
	fmt.Println("I am from foo")

}

// 1 param no returns
func bar(s string) {
	fmt.Println("My name is", s)
}

// 1 param, 1 return
func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

// 2 params, 2 returns
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years"), age
}

// variadic parameters
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}
