package crypto

import (
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

func (h *Hash) Generate(s string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	pepper := string(65 + rand.Intn(26))
	pepperBytes := []byte(s + pepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pepperBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes[:]), nil
}

func (h *Hash) Compare(hash string, p string) error {
	for i := 0; i < 26; i++ {
		pepperBytes := []byte(p + string(65+i))
		err := bcrypt.CompareHashAndPassword([]byte(hash), pepperBytes)
		if err == nil {
			return nil
		}
	}
	return errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
}
