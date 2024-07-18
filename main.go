package main

import "log"

func main() {
	BPass, err := BcryptHash()
	if err != nil {
		log.Println("error in generating the hashed password from bcrypt")
		return
	}

	log.Printf("Bcrypt password is : %s \n", BPass)
	ShaPass := ShaHash()
	log.Printf("Sha256 password is : %v \n", ShaPass)

	// ****************conclusion is************************

	/*
		sha256 produce the same hashed password every time for the same string.
		bcrypt produce different hashed password for the same string everytime.
	*/

}
