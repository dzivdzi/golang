package main

import (
	"fmt"
)

type Person struct {
	FirstName       string
	LastName        string
	FavIceCreamFlav string
}

type Engine struct {
	Electric bool
}

type Vehicle struct {
	Engine
	Make  string
	Model string
	Doors int
	Color string
}

func main() {
	// 1
	people := []Person{
		{
			"Glo",
			"Stoj",
			"Vanilla",
		},
		{
			"Alex",
			"Stojk",
			"Banana",
		},
	}

	for _, v := range people {
		fmt.Printf("%s's fav ice-cream flavour is %s\n", v.FirstName, v.FavIceCreamFlav)
	}

	// 2
	mappedPeople := map[string]string{}
	for _, v := range people {
		fmt.Println(v)
		mappedPeople[v.LastName] = v.FavIceCreamFlav
	}

	fmt.Println(mappedPeople)

	// 3
	cars := []Vehicle{
		{
			Engine: Engine{true},
			Make:   "Toyota",
			Model:  "Yaris",
			Doors:  4,
			Color:  "Blue",
		},
		{
			Engine{false},
			"Toyota",
			"Rav4",
			5,
			"Green",
		},
	}

	for _, v := range cars {
		fmt.Printf("%s \t %v\n", v.Model, v.Engine)
	}

	// 4
	persona := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Alex",
		friends: map[string]int{
			"Robi":   27,
			"Babuka": 27,
		},
		favDrinks: []string{"Beer", "Lemonade"},
	}
	fmt.Println(persona)
	fmt.Printf("%T\n", persona)
}
