package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Port   string
	Router *mux.Router
}

func NewServer(port string) *Server {
	s := Server{port, mux.NewRouter()}
	configureRoutes(&s)
	return &s
}

func (s *Server) Start() error {
	log.Println("Starting server on port", s.Port)
	return http.ListenAndServe(s.Port, s.Router)
}
