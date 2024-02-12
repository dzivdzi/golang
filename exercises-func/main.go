package main

import (
	"fmt"
	"math"
)

type User struct {
	first string
	age   int
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) Area() float64 {
	return s.length * s.width
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)

}

type shape interface {
	Area() float64
}

func info(s shape) float64 {
	return s.Area()
}

func main() {

	sq := square{
		length: 10.0,
		width:  10.0,
	}

	c := circle{
		radius: 10,
	}

	fmt.Println(info(sq))
	fmt.Println(info(c))

	/*
		foo := foo()
		i, s := bar()
		fmt.Println(foo, i, s)
	*/
	/*
	   xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	   fmt.Println(foo(xi...))

	   defer iAmDone()
	   sleepyTime()
	*/

	/*
		u := User{
			first: "John",
			age:   23,
		}

		u.speak()
	*/

}

/*
//4
func (u User) speak() {
	fmt.Printf("My name is %s and I am %v years old.\n", u.first, u.age)
}
*/
/*
// 1
func foo() int {
	return 2
}
func bar() (int, string) {
	return 3, "nice"
}
*/
// 2
/*
func foo(ii ...int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func bar(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
*/

// 3
/*
func iAmDone() {
	fmt.Println("I am done")
}

func sleepyTime() {
	time.Sleep(5 * time.Second)

	fmt.Println("Sleep is over")
}
*/
/*
// named func
func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}

	return
}
*/
