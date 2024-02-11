package main

import (
	"fmt"
)

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was ", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 26
	an["Stephanie"] = 24
	an["George"] = 78
	fmt.Printf("%#v\n", an)

	for k, v := range an {
		fmt.Println(k, v)
	}

	for k := range an {
		fmt.Println(k)
	}
	for _, v := range an {
		fmt.Println(v)
	}

	// Delete an element from a map
	delete(an, "George")

	// Comma OK Idiom
	v, ok := an["Lucas"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Doesn't exist")
	}
}
