package server

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/application"
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

type WeightedRoundRobinLoadBalancer struct {
	discovery       Discovery
	lastServedIndex int
	//Mutex for locking NextServer
	mutex sync.Mutex
}

func SelectLoadBalancer(e application.Environment) LoadBalancer {
	return newRoundRobinLoadBalancer(e)
}

func newRoundRobinLoadBalancer(e application.Environment) *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		discovery:       NewConfigBasedDiscovery(e.ConfigPath),
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

//
//func (lb *WeightedRoundRobinLoadBalancer) NextServer() *Server {
//	//TODO random number max to sum of weights
//	//Calculate boundaries and return
//}
