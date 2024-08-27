package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	userPassword1 := "some user-provided password"

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	// Store this "hash" somewhere, e.g. in your database

	// After a while, the user wants to log in and you need to check the password he entered
	userPassword2 := "some user-provided password"
	hashFromDatabase := hash

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Password was correct!")
}
