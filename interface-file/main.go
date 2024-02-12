package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// A byte buffer is a region of memory used to temporarily store a sesquence of bytes. It provides a data structure for efficient manipulation of bytedata
// A byte buffer allows you to read and write bytes to and from the buffer making it useful for tasks like data serialization, network communication, file I/o and efficient string manipulation.

// The concept of a byte buffer is not specific to any particular programmin language  but is a general concept in computer programming.
// Different programming languages may provide their own implementations of byute buffers each with it's own set of features and functionality

type Person struct {
	first string
}

// This takes a variable which takes the writer interface
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p)) 
// and any error encountered that caused the write to stop early. 
// Write must return a non-nil error if it returns n < len(p). 
// Write must not modify the slice data, even temporarily.
func (p Person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))

}

func main() {
	p := Person{
		first: "Jenny",
	}
	// Creates a file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	// Defers closing the file
	defer f.Close()

	// Create a var of type Buffer
	var c bytes.Buffer

	// Write to the file created above
	p.writeOut(f)
	// In the write out function, it takes a person of name "p" (this is what we called the variable) and writes it into the buffer
	// Note that we have to put a pointer to the buffer address in order to work as we can't modify a copy
	// In this case, the buffer created is called "c" (check above) so what this will do is
	// take the var "p" from above, run the function write out and write to the buffer the first name of the person (also defined above)
	p.writeOut(&c)
	// Shows what is inside the buffer
	fmt.Println(c.String())

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	// .WriteString appends an additional string to what is already stored in "b"
	b.WriteString("Gophers")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("It's Sunday")
	fmt.Println(b.String())
}
