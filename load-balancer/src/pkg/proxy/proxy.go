package proxy

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func SimpleBalancedProxy() *BalancedProxy {
	return new(BalancedProxy)
}

type BalancedProxy struct {
}

func (p *BalancedProxy) RegisterRoutes(s *server.Server) {
	s.Router.HandleFunc("/", proxyRequest(p))
}

func proxyRequest(p *BalancedProxy) http.HandlerFunc {
	server := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(server)

	return func(w http.ResponseWriter, r *http.Request) {
		rProxy.ServeHTTP(w, r)
	}
}

func getServer() *url.URL {
	parsedUrl, _ := url.Parse("http://127.0.0.1:8081")
	return parsedUrl
}
