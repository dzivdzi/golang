package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}() // the last set of () are called executing brackets - these tell the compiler that the func needs to get executed

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Alex") // if we want the function to take something, we put the value in the second brackets, and we place the type we want it to take in the first ones

	// func expression - assigning a func to a variable
	// an expression is a combination of values, vatiables, operators and function calls that are
	// evaluated to produce a single value

	x := func() {
		fmt.Println("I am now x, not anonymous")
	}

	y := func(s string) {
		fmt.Println("I am not anonymous anymore - your name is: ", s)
	}

	x()
	y("Alex")

	z := doMath(5, 5, add)
	n := doMath(10, 5, subtract)
	fmt.Println(z)
	fmt.Println(n)

	f := incrementor() // initial value is 0 and we get a function f() which increments only that value inside according to the scope of the function Incrementor
	fmt.Println(f())   // call function f (f is now a function which just increments the value of +1) - now we get 1, we don't get 0 + 1 but just 1 as x is enclosed in the first incrementor function
	fmt.Println(f())   // again we call f - now f works on value 1 as it became 1 prev now it's gonna be 2
	fmt.Println(f())   //3
	fmt.Println(f())   //4
	fmt.Println(f())   //5
	fmt.Println(f())   //6

	fmt.Println(factorial(5))

}

func foo() {
	fmt.Println("foo ran")
}

// passing a func as an argument - callback
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

// This function takes 2 ints and a function which takes two ints and returns an int - it returns an int
// As long as the function which is passed as the intake parameter conforms to the function which is required, it will run
// Meaning if the function which is passed as an entry parameter takes 2 ints and returns an int, and the original function wants two ints and returns an int, it will work
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// Closure
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Recursion
func factorial(n int) int {
	// Func is for looping itself until the value becomes 0 as in teh if clause
	// We can also do this in a for loop

	fmt.Println("N is:", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
