package proxy

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/application"
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"log"
	"net/http"
)

func HelloBalancedProxy(e application.Environment) *HelloProxy {
	return &HelloProxy{
		lb: server.SelectLoadBalancer(e),
	}
}

type HelloProxy struct {
	lb server.LoadBalancer
}

func (p *HelloProxy) RegisterRoutes(s *application.Application) {
	log.Println("Registering proxy on /hello route")
	s.Router.HandleFunc("/hello", proxyRequest(p))
}

func proxyRequest(p *HelloProxy) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("Proxying request:", req.Method, req.RequestURI, req.Body)
		server := p.lb.NextServer()
		log.Println("Selected server:", server)
		server.ServeRequest(rw, req)
	}
}
