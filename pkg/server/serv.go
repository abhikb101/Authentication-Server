package server

import (
	"log"
	"net/http"
	"os"
	root "user_auth/pkg"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(u root.UserService) *Server {
	s := Server{router: mux.NewRouter()}
	NewRegisterRouter(u, s.NewRoute("/newuser"))
	return &s
}

func (s *Server) Start() {
	log.Printf("Listening on port 12345")
	if err := http.ListenAndServe(":12345", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("Error")
	}

}

func (s *Server) NewRoute(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
