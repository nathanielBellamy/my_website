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
  pw := os.Getenv("PW")
  pw_hash, err := h.Generate(pw)
  if err != nil {
    return false
  }
  res := bcrypt.CompareHashAndPassword([]byte(pw_hash), incomingPw)
  return res == nil
}
