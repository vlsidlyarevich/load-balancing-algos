package server

import (
	"github.com/gorilla/mux"
	"github.com/vlsidlyarevich/load-balancer/pkg/proxy"
	"log"
	"net/http"
)

type Server struct {
	Port   string
	Router *mux.Router
}

func NewServer(port string) *Server {
	s := Server{port, mux.NewRouter()}
	return &s
}

func (s *Server) Start() error {
	log.Println("Starting server on port", s.Port)
	return http.ListenAndServe(s.Port, s.Router)
}

func (s *Server) AddProxy(p proxy.Proxy) {
	p.RegisterRoutes(s)
}
