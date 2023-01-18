package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password12453`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	loginpass := `passwor12453`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpass))
	if err != nil {
		fmt.Println("Cant Login")
	}
	fmt.Println("You are logged in")
}
