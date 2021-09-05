package application

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RouteHandler interface {
	RegisterRoutes(s *Application)
}

type Application struct {
	Env    Environment
	Port   string
	Router *mux.Router
}

type Environment struct {
	Args       []string
	ConfigPath string
}

func NewApplication(port string, args []string, configPath string) *Application {
	s := Application{Environment{Args: args, ConfigPath: configPath}, port, mux.NewRouter()}
	return &s
}

func (s *Application) Start() error {
	log.Println("Starting application server on port", s.Port)
	return http.ListenAndServe(s.Port, s.Router)
}

func (s *Application) RouteHandler(h RouteHandler) {
	log.Printf("Registering RouteHandler %+v\n", h)
	h.RegisterRoutes(s)
}
