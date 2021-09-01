package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

const DefaultConfigPath = "conf/config.toml"

func main() {
	var args = os.Args[1:]
	log.Println("Arguments used:", args)

	var configPath = DefaultConfigPath
	if len(args) > 0 && args[0] != "" {
		configPath = args[0]
	}
	var config = LoadConfig(configPath)

	var server = NewServer(":"+config.Server.Port, config.Server.Id)
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
