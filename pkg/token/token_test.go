package token_test

import (
	"testing"
	root "user_auth/pkg"
	"user_auth/pkg/token"
)

const (
	testUsername = "integration_testser"
	testEmail    = "integration_test_email"
	testPassword = "integration_test_pword"
)

func Test_JwtService(t *testing.T) {
	t.Run("GenerateJwt", GenerateAccessToken_should_generate_jwt)
}

func GenerateAccessToken_should_generate_jwt(t *testing.T) {
	//Arrange
	j := token.TokenService{}
	user := root.User{
		Email:    testEmail,
		Username: testUsername,
		Password: testPassword,
	}
	token, err := j.GenerateAccessToken(&user)
	token1 := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImludGVncmF0aW9uX3Rlc3RfZW1haWwiLCJleHAiOjE1ODg0MzUzNzUsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJ1c2VybmFtZSI6ImludGVncmF0aW9uX3Rlc3RzZXIifQ.hYeFhnXxm2SAf3g52erGl3Shc5xrXFXNcpJNjOeSv5OKmfM2EtDcxuFe43TnVIPifHdXmYNNGeFw9O_oobUUh82LNZWAvfMZFLuEKuYZa2we_tIlHpsenBxfXFKaBxf1py6dMv2GtjIHMZSmLitTgwh_h7hRBiTUEKv4E7ErlfquBmFtvVwpRFS7ljll2k6Amiqd7s6iNcnbmXQmScb-uHhExacP4IldVpCiDYh33RBaWizfbTiJcq_kc0wr-_g6Uoa7n5zarEkuQe3fzzOotA1awl08xp-MNfwvZd4TDIR5Sq88XBSKp7guhOSG4NnnUbXwDJYyCmhAL2NAi2ZoYQ"
	token2 := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImludGVncmF0aW9uX3Rlc3RfZW1haWwiLCJleHAiOjE1ODg0MzUzNzUsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJ1c2VybmFtZSI6ImludGVncmF0aW9uX3Rlc3RzZXIifQ.hYeFhnXxm2SAf3g52erGl3Shc5xrXFXNcpJNjOeSv5OKmfM2EtDcxuFe43TnVIPifHdXmYNNGeFw9O_oobUUh82LNZWAvfMZFLuEKuYZa2we_tIlHpsenBxfXFKaBxf1py6dMv2GtjIHMZSmLitTgwh_h7hRBiTUEKv4E7ErlfquBmFtvVwpRFS7ljll2k6Amiqd7s6iNcnbmXQmScb-uHhExacP4IldVpCiDYh33RBaWizfbTiJcq_kc0wr-_g6Uoa7n5zarEkuQe3fzzOotA1awl08xp-MNfwvZd4TDIR5Sq88XBSKp7guhOSG4NnnUbXwDJYyCmhAL2NAi2ZoYQ"
	token3 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImludGVncmF0aW9uX3Rlc3RfZW1haWwiLCJleHAiOjE1ODg0MzUzNzUsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJ1c2VybmFtZSI6ImludGVncmF0aW9uX3Rlc3RzZXIifQ.hYeFhnXxm2SAf3g52erGl3Shc5xrXFXNcpJNjOeSv5OKmfM2EtDcxuFe43TnVIPifHdXmYNNGeFw9O_oobUUh82LNZWAvfMZFLuEKuYZa2we_tIlHpsenBxfXFKaBxf1py6dMv2GtjIHMZSmLitTgwh_h7hRBiTUEKv4E7ErlfquBmFtvVwpRFS7ljll2k6Amiqd7s6iNcnbmXQmScb-uHhExacP4IldVpCiDYh33RBaWizfbTiJcq_kc0wr-_g6Uoa7n5zarEkuQe3fzzOotA1awl08xp-MNfwvZd4TDIR5Sq88XBSKp7guhOSG4NnnUbXwDJYyCmhAL2NAi2ZoYQ"

	if err != nil {
		t.Error(err)
	}
	t.Log("\n")
	t.Log("\n")
	t.Log(token)
	t.Log("\n")
	t.Log("\n")
	err = j.VerifyToken(token)
	if err != nil {
		t.Log("FAILED")
		t.Error(err)
	}
	err = j.VerifyToken(token1)
	t.Log(err)
	err = j.VerifyToken(token2)
	t.Log(err)
	err = j.VerifyToken(token3)
	if err != nil {
		t.Log(err)
	}

}
