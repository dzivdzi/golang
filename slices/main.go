package main

import "fmt"

func main() {
	xs := []string{"Something", "Nice"}
	fmt.Println(xs)

	fmt.Printf("%T\n", xs)

	flavours := []string{"Almond Biscotti Café", "Banana Pudding", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)",
		"Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ",
		"Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	fmt.Println(flavours)
	fmt.Println(len(flavours))

	for i, v := range flavours {
		fmt.Printf("%v: %s\n", i+1, v)
	}
	// variadic parameter (...) - means you can add as many values from the same type as you want
	xs = append(xs, "This", "is")
	fmt.Println(xs)
	fmt.Println("---------------------------")

	// Slicing
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// [inclusive:exclusive]
	fmt.Printf("numbers %#v\n", numbers[0:4])
	fmt.Println("---------------------------")
	// [:exclusive]
	fmt.Printf("numbers %#v\n", numbers[:6])
	fmt.Println("---------------------------")
	// [inclusive:]
	fmt.Printf("numbers %#v\n", numbers[3:])
	fmt.Println("---------------------------")

	// Deleting from a slice
	numbers = append(numbers[:4], numbers[5:]...)
	fmt.Println("---------------------------")
	fmt.Println(numbers)

	// Composite Literal
	si := []string{"a", "b", "c"}

	// Make
	xi := make([]int, 0, 10)
	fmt.Println("---------------------------")
	fmt.Println(si)
	fmt.Println("---------------------------")
	fmt.Println(cap(xi))

	// Multidimensional slices
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	mp := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}

	xxs := [][]string{jb, mp}
	fmt.Println("---------------------------")
	fmt.Println(xxs)
}
