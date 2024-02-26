package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("newFile", []byte("1\n\n\n123"), 0744)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}
