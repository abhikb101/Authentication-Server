package server

import (
	"fmt"
	"net/http"
	root "user_auth/pkg"

	"github.com/gorilla/mux"
)

type RegisterRouter struct {
	user root.UserService
}

func NewRegisterRouter(u root.UserService, s *mux.Router) *mux.Router {
	registerrouter := RegisterRouter{u}
	s.HandleFunc("/Register", registerrouter.RegisterHandler)
	return s
}

func (router *RegisterRouter) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Register")
}
