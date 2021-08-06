package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	//TODO fetch name from config or generate unique
	var server = NewServer(":8081", "Server-1")
	log.Fatal(server.Start())
}

type Server struct {
	Port   string
	Name   string
	Router *mux.Router
}

func NewServer(port string, name string) *Server {
	s := Server{port, name, mux.NewRouter()}
	ConfigureRoutes(&s)
	return &s
}

func (s *Server) Start() error {
	log.Println("Starting server on port", s.Port)
	return http.ListenAndServe(s.Port, s.Router)
}

func ConfigureRoutes(s *Server) {
	s.Router.HandleFunc("/hello", HelloHandler(s.Name)).Methods("GET")
}

func HelloHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, name)
		if err != nil {
			log.Println("Error during handling response: ", err)
		}
	}
}
