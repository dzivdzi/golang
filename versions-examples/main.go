package main

import (
	"fmt"

	"github.com/dzivdzi/puppy"
)

func main() {
	fmt.Println("Working on versioning!")

	var blankString string
	blankString = "I am a blank string"

	notABlankString := "I am not a blank string :)"

	fmt.Println(puppy.LittleDog)
	fmt.Println(puppy.BigDog)
	fmt.Println(blankString)
	fmt.Println(notABlankString)
}
