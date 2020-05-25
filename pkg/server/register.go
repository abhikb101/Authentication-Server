package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	root "user_auth/pkg"

	"github.com/gorilla/mux"
)

type RegisterRouter struct {
	user  root.UserService
	token root.TokenService
}

func NewRegisterRouter(u root.UserService, t root.TokenService, s *mux.Router) *mux.Router {
	registerrouter := RegisterRouter{u, t}
	s.HandleFunc("/Register", registerrouter.RegisterFormHandler).Methods("GET").Name("Register_Form")
	s.HandleFunc("/Register", registerrouter.RegisterHandler).Methods("POST").Name("Register_Function")
	return s
}

func (router *RegisterRouter) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user root.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = router.user.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := router.token.GenerateAccessToken(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, token)
}

func (router *RegisterRouter) RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Query())
}
