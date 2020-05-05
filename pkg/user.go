package root

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
