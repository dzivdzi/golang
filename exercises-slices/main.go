package main

import "fmt"

func main() {
	// 1
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("Index: %v \t Value: %v \t Type: %T\n", i, v, v)
	}
	// 2
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range b {
		fmt.Printf("Index: %v \t Value: %v\n", i, v)
	}
	// 3
	fmt.Println(b[:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])
	// 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	// 5
	x = append(x[:4], x[8:]...)
	fmt.Println(x)
	// 6
	states := make([]string, 0, 50)
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(states), cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("%s \t %v\n", states[i], i)
	}
	// 7
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "I'm 008."}
	jbm := [][]string{jb, mp}
	for i, v := range jbm {
		fmt.Printf("Ranging over %v slice\n", i+1)
		for _, k := range v {
			fmt.Printf("%s\n", k)
		}
	}
}
