package main

import "fmt"

type user struct {
	first string
}

func changeFirst(u user, s string) user {
	u.first = s
	return u
}

func changeFirstP(u *user, s string) {
	u.first = s
}

func main() {
	p := user{
		first: "Jenny",
	}
	fmt.Println(p)
	p = changeFirst(p, "Bob")
	fmt.Println(p)
	changeFirstP(&p, "mara")
	fmt.Println(p)
}
