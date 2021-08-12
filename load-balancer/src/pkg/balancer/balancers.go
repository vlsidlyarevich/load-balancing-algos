package balancer

type LoadBalancer interface {
	NextServer() *Server
}

type RoundRobinLoadBalancer struct {
	servers []*Server
	//TODO
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		servers: getServers(),
	}
}

func (lb *RoundRobinLoadBalancer) NextServer() *Server {
	//TODO
	panic("implement me")
}

func getServers() []*Server {
	return []*Server{
		newServer("http://127.0.0.1:8081"),
		newServer("http://127.0.0.1:8082"),
		newServer("http://127.0.0.1:8083"),
	}
}
