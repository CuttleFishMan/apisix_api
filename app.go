package apisix

import "time"

type Svc struct {
	Name    string
	Port    string
	XAPIKEY string

	// Default eth0 addr
	InetIp string

	// APISIX host
	// Default 127.0.0.1:9180
	Host    string
	Version string

	// Default ["127.0.0.1"]
	Hosts []string

	// Default ["0.0.0.0/0"]
	Remote_Addrs []string

	// Default false
	EnableWebsocket bool

	// Default ["PUT", "GET","POST","PATCH","DELETE","OPTIONS","HEAD","CONNECT","TRACE"],
	Methods []string
	HTTP2   bool

	Plugins map[string]interface{}
}

func (s *Svc) RegisterService() error {
	return s.registerService()
}

// query routes and update nodes
func (s *Svc) RegisterRouter(router string, ttls ...time.Duration) error {
	return s.registerRouter(router, ttls...)
}
