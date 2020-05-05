package token

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"
	root "user_auth/pkg"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenService struct {
}

func (j *TokenService) GenerateAccessToken(u *root.User) (string, error) {
	rawkey, err := ioutil.ReadFile("C:/Users/HP-PC/go/src/user_auth/keys/private.key")
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(rawkey)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["type"] = "access_token"
	claims["authorized"] = true
	claims["username"] = u.Username
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (j *TokenService) GenerateRefreshToken(u *root.User) (string, error) {
	rawkey, err := ioutil.ReadFile("C:/Users/HP-PC/go/src/user_auth/keys/private.key")
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(rawkey)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["type"] = "refresh_token"
	claims["authorized"] = true
	claims["username"] = u.Username
	claims["email"] = u.Email
	claims["userid"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 5).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (j *TokenService) GenerateAuthorizationCode(u *root.User) (string, error) {
	token := make([]byte, 32)
	rand.Read(token)
	return fmt.Sprintf("%x", token), nil
}

func (j *TokenService) VerifyToken(t string) error {
	rawkey, err := ioutil.ReadFile("C:/Users/HP-PC/go/src/user_auth/keys/public.key")
	if err != nil {
		return err
	}
	block, _ := pem.Decode(rawkey)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	_, err = jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Header["alg"] == "RS256" {
			return key, nil
		}
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	})
	if err != nil {
		return err
	}

	return nil
}
