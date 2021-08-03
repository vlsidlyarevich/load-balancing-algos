package server

import "github.com/vlsidlyarevich/load-balancer/pkg/handler"

func configureRoutes(s *Server) {
	s.Router.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
}
