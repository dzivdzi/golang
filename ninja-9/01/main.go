package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hey from a go one")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hey from a go two")
		wg.Done()
	}()
	fmt.Println("About to exit")

	wg.Wait()
}
