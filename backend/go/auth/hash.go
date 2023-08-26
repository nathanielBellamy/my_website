package auth

import (
	"os"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{
  Log *zerolog.Logger
}

//Generate a salted hash for the input string
func (h *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
    h.Log.Error().
          Err(err).
          Msg("Generate Client PW Hash ERROR")
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
    h.Log.Error().
          Err(err).
          Msg("Generate Server PW Hash ERROR")
    return false
  }
  res := bcrypt.CompareHashAndPassword([]byte(pw_hash), incomingPw)
  return res == nil
}
