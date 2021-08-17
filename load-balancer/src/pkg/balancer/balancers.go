package balancer

import (
	"sync"
)

type LoadBalancer interface {
	NextServer() *Server
}

type RoundRobinLoadBalancer struct {
	servers         []*Server
	lastServedIndex int
	//Mutex for locking NextServer
	mutex sync.Mutex
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		servers:         getServers(),
		lastServedIndex: -1,
	}
}

func (lb *RoundRobinLoadBalancer) NextServer() *Server {
	lb.mutex.Lock()

	if lb.lastServedIndex > len(lb.servers) {
		lb.lastServedIndex = 0
	}
	server := lb.servers[lb.lastServedIndex]
	lb.lastServedIndex++

	defer lb.mutex.Unlock()

	return server
}

func getServers() []*Server {
	return []*Server{
		newServer("http://127.0.0.1:8081"),
		newServer("http://127.0.0.1:8082"),
		newServer("http://127.0.0.1:8083"),
	}
}
