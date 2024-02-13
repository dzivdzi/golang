// Method Sets
package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) Walk() {
	fmt.Println("My name is", d.first, "and I'm walking!")
}

func (d *Dog) Run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type Youngin interface {
	Walk()
	Run()
}

func YoungRun(y Youngin) {
	y.Run()
}

func main() {
	d1 := Dog{"Henry"}
	d1.Walk()
	d1.Run()
	YoungRun(&d1)
	d2 := &Dog{"Padget"}
	d2.Walk()
	d2.Run()
	YoungRun(d2)

	
}
