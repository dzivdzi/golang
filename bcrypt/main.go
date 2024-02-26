package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(string(bs))
}
