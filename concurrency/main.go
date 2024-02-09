package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// CONDITIONAL
	// if statements
	// switch statements

	// conucurrency
	// select a channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Microsecond)
		ch1 <- 41 // 41 is pushed to ch1 (we can choose the number if we wanted to)
	}()
	go func() {
		time.Sleep(d2 * time.Microsecond)
		ch2 <- 42 // 42 is pushed to ch1 (we can choose the number if we wanted to)
	}()

	// A "select" statement chooses which of a set of possible send or recive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations

	select { // This case, select statement chooses the first one to finish or in other words, select the first function that finishes, and print out the value which was sent to the respective channel
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}
}
