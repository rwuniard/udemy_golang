package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Hello")

	pwd := `password123`
	// Notes: when you run this over and over again, it won't generate the same encrypted password.
	// You need to use CompareHashAndPassword function to compare password.
	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(encryptedPwd))

	// Comparing password
	loginPwd := `password123`
	err = bcrypt.CompareHashAndPassword(encryptedPwd, []byte(loginPwd))
	if err != nil {
		fmt.Println("Password doesn't match")
		return
	}

	fmt.Println("Password matched")
}
