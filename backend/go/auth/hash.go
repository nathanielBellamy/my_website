package auth

import (
	"golang.org/x/crypto/bcrypt"
  "os"
)

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
func (h *Hash) Compare(incoming string) bool {
  incomingPw := []byte(incoming)
  pw_hash := os.Getenv("PW_HASH")
  res := bcrypt.CompareHashAndPassword([]byte(pw_hash), incomingPw)
  return res == nil
}
