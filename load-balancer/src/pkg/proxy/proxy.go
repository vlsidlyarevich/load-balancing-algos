package proxy

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func HelloBalancedProxy() *HelloProxy {
	return new(HelloProxy)
}

type HelloProxy struct {
}

func (p *HelloProxy) RegisterRoutes(s *server.Server) {
	log.Println("Registering proxy on /hello route")
	s.Router.HandleFunc("/hello", proxyRequest(p))
}

func proxyRequest(p *HelloProxy) http.HandlerFunc {
	server := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(server)

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Proxying request", r.Method, r.RequestURI, r.Body)
		log.Println("Selected server", server)
		rProxy.ServeHTTP(w, r)
	}
}

func getServer() *url.URL {
	parsedUrl, _ := url.Parse("http://127.0.0.1:8081")
	return parsedUrl
}
