package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	root "user_auth/pkg"

	"github.com/gorilla/mux"
)

type LoginRouter struct {
	user  root.UserService
	token root.TokenService
}

func NewLoginRouter(u root.UserService, t root.TokenService, s *mux.Router) *mux.Router {
	loginrouter := LoginRouter{u, t}
	s.HandleFunc("/Login", loginrouter.LoginHandler).Methods("GET").Name("Login")
	s.PathPrefix("/api").Subrouter().HandleFunc("/Token", loginrouter.TokenHandler).Methods("GET").Name("token")
	return s
}

func (router *LoginRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user root.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	us, err := router.user.Login(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := router.token.GenerateRefreshToken(us)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, token)
}
func (router *LoginRouter) TokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Values("Auth")
	accesstoken, err := router.token.RefreshTokenExchange(token[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, accesstoken)
}
