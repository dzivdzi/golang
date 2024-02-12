package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// Wrapper function (wrapper) is a function that provides an additional layer around an existing function or method.
// It acts as an intermediary between the caller and the wrapped function allowing to modify inputs oputputs or behaviour without directly mmodifying the original function
// Below func:
// logInfo - name (given by the creator)
// Takes a value(s)(string) which is of type fmt.Stringer(given by a package - fmt package)
// performs a printline (part of another package) with an additional input provided by the creator and calles the method .String as the function .String is present on the Stringer type (which is an interface in the fmt package)
// In our case both of our defined TYPES have the String Method and are created to be part of the STRINGER interface which is part of the fmt package
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}
func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}
