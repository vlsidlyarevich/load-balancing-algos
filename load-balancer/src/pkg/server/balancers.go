package server

import (
	"sync"
)

type LoadBalancer interface {
	NextServer() *Server
}

type RoundRobinLoadBalancer struct {
	discovery       Discovery
	lastServedIndex int
	//Mutex for locking NextServer
	mutex sync.Mutex
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		discovery:       NewConfigBasedDiscovery(),
		lastServedIndex: 0,
	}
}

func (lb *RoundRobinLoadBalancer) NextServer() *Server {
	lb.mutex.Lock()
	servers := lb.discovery.ServerList()

	if lb.lastServedIndex >= len(servers) {
		lb.lastServedIndex = 0
	}
	server := servers[lb.lastServedIndex]
	lb.lastServedIndex++

	defer lb.mutex.Unlock()

	return server
}
