package proxy

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/application"
	"github.com/vlsidlyarevich/load-balancer/pkg/balancer"
	"log"
	"net/http"
)

func HelloBalancedProxy() *HelloProxy {
	return &HelloProxy{
		lb: balancer.NewRoundRobinLoadBalancer(),
	}
}

type HelloProxy struct {
	lb balancer.LoadBalancer
}

func (p *HelloProxy) RegisterRoutes(s *application.Application) {
	log.Println("Registering proxy on /hello route")
	s.Router.HandleFunc("/hello", proxyRequest(p))
}

func proxyRequest(p *HelloProxy) http.HandlerFunc {
	server := p.lb.NextServer()

	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("Proxying request", req.Method, req.RequestURI, req.Body)
		log.Println("Selected application", server)
		server.ServeRequest(rw, req)
	}
}
