package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	url     *url.URL
	isAlive bool
	rProxy  *httputil.ReverseProxy
}

func (s Server) String() string {
	return fmt.Sprintf("Server url: %s, isAlive: %t", s.url, s.isAlive)
}

func (s Server) ServeRequest(rw http.ResponseWriter, req *http.Request) {
	s.rProxy.ServeHTTP(rw, req)
}

func newServer(rawurl string) *Server {
	parsedUrl, _ := url.Parse(rawurl)

	return &Server{
		url:     parsedUrl,
		isAlive: true,
		rProxy:  httputil.NewSingleHostReverseProxy(parsedUrl),
	}
}
