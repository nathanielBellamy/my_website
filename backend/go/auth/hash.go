package dev_auth

import (
	// "errors"
	// "strings"

	"golang.org/x/crypto/bcrypt"
)

//Hash implements root.Hash
type Hash struct{}

//Generate a salted hash for the input string
func (h *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (h *Hash) Compare(existingHash []byte, incoming string) bool {
  incomingPw := []byte(incoming)
  res := bcrypt.CompareHashAndPassword(existingHash, incomingPw)
  return res == nil
}
