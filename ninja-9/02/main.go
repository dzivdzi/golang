package main

import "fmt"

type Person struct {
}

func (p *Person) speak() {
	fmt.Println("I am human")
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p1 := Person{}

	saySomething(&p1)
}
