package root

import "github.com/gorilla/mux"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
	Login(u *User) (*User, error)
	GetUsers() ([]User, error)
}

type Hash interface {
	Generate(s string) (string, error)
	Compare(hash string, p string) error
}

type TokenService interface {
	GenerateAccessToken(u *User) (string, error)
	GenerateRefreshToken(u *User) (string, error)
	GenerateAuthorizationCode(u *User) (string, error)
	VerifyToken(t string) error
}

type server interface {
	Start()
	NewRoute(path string) *mux.Router
}
