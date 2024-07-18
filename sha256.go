package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func ShaHash() string {
	s := "navneet shukla"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	hashString := hex.EncodeToString(bs)

	return hashString
}
