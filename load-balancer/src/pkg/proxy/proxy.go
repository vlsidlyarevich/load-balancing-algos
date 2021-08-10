package proxy

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/clients"
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"io"
	"log"
	"net/http"
)

type Proxy interface {
	RegisterRoutes(s *server.Server)
}

func NewHelloProxy() *HelloProxy {

	return new(HelloProxy)
}

type HelloProxy struct {
	Client *clients.EchoClient
}

func proxyHello(p *HelloProxy) http.HandlerFunc {

	var response = p.Client.ForwardEcho()

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, response)
		if err != nil {
			log.Println("Error during handling response: ", err)
		}
	}
}

func (p *HelloProxy) RegisterRoutes(s *server.Server) {
	s.Router.HandleFunc("/hello", proxyHello(p)).Methods(http.MethodGet)
}
