package main

import "fmt"

type Person struct {
	first string
}

type SecretAgent struct {
	Person
	Ltk bool
}

func (p Person) speak() {
	fmt.Println("I am", p.first)
}

func (p SecretAgent) speak() {
	fmt.Println("I am secret agent", p.first)
}

type Human interface {
	speak()
}

// Every single type which has the method speak can fall in the category Human
// In our case, both Person type and SecretAgent type have the method called speak
// This allows the function below to take more than one type and perform it's actions to
// A value can be of more than one type - that's polymorphism
func saySomething(h Human) {
	h.speak()
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			first: "James",
		},
		Ltk: true,
	}
	p2 := Person{
		first: "Jenny",
	}

	saySomething(sa1)
	saySomething(p2)

	sa1.speak()
	p2.speak()
}
