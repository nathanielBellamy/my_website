package auth

import (
  "fmt"
	"golang.org/x/crypto/bcrypt"
  "os"
)

type Hash struct{}

//Generate a salted hash for the input string
func (h *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
    fmt.Printf("ERR Generate Hash ERROR")
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (h *Hash) Compare(incoming string) bool {
  incomingPw := []byte(incoming)
  pw_hash := os.Getenv("PW_HASH")
  fmt.Printf("\n pw_hash from os: %v \n", pw_hash)
  pw_hash_manual := "$2a$10$eoka2klp4SoOA4mXyiHkQuctdKkXXJfalLotfvX7hbuiryu5fQA.G"
  fmt.Printf("\n pw_hash manual: %v \n", pw_hash_manual)
  res := bcrypt.CompareHashAndPassword([]byte(pw_hash_manual), incomingPw)
  return res == nil
}
