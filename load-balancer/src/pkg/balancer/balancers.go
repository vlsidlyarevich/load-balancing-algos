package balancer

import (
	"sync"
)

type LoadBalancer interface {
	NextServer() *Server
}

type RoundRobinLoadBalancer struct {
	servers         []*Server
	lastServedIndex int16
	//Mutex for locking NextServer
	mutex sync.Mutex
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		servers: getServers(),
	}
}

func (lb *RoundRobinLoadBalancer) NextServer() *Server {
	lb.mutex.Lock()

	defer lb.mutex.Unlock()
}

func getServers() []*Server {
	return []*Server{
		newServer("http://127.0.0.1:8081"),
		newServer("http://127.0.0.1:8082"),
		newServer("http://127.0.0.1:8083"),
	}
}
