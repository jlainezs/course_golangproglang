package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "ThisIsAPassword"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(bs)

	loginAttemptPwd := "ThisIsAPassword"
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginAttemptPwd))
	if err != nil {
		panic(err)
	}

	loginAttemptPwd = "ThisIsAPasswordAAAAAA"
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginAttemptPwd))
	if err != nil {
		panic(err)
	}
}
