package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	var server = NewServer(":8081")
	log.Fatal(server.Start())
}

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

func configureRoutes(s *Server) {
	s.Router.HandleFunc("/hello", HelloHandler).Methods("GET")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello, world!\n")
	if err != nil {
		log.Println("Error during handling response: ", err)
	}
}
