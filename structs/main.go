package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   37,
	}

	people := []person{}

	people = append(people, p1, p2)
	fmt.Println(people)

	// embedded structs

	p3 := secretAgent{
		person: person{
			first: "Alex",
			last:  "Stojk",
			age:   33,
		},
		ltk: true,
	}

	fmt.Println(p3)
	fmt.Println(p3.person)

	// anonymous structs

	pers := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   35,
	}

	fmt.Printf("%T\n", pers)
	fmt.Printf("%T\n", p1)

	// Composition

	type foo int
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
