package main

import "golang.org/x/crypto/bcrypt"

func BcryptHash() (string, error) {
	password := []byte("navneet shukla")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	} else {
		return string(hashedPassword), nil
	}
}
