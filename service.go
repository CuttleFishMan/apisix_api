package apisix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/eavesmy/golang-lib/net"
)

const (
	UPSTREAM_URI = "/admin/upstreams"
	SERVICE_URI  = "/admin/services"
	ROUTER_URI   = "/admin/routes"
)

/*
curl http://127.0.0.1:9180/apisix/admin/routes/apisix.eva7base.com -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -i -d '
{
    "uri": "/*",
    "hosts": ["apisix.eva7base.com"],
    "remote_addrs": ["0.0.0.0/0"],
    "methods": ["PUT", "GET","POST","PATCH","DELETE","OPTIONS","HEAD","CONNECT","TRACE"],
    "enable_websocket": false,
    "upstream": {
        "type": "roundrobin",
        "nodes": {
            "127.0.0.1:9000": 1
        }
    }
}'
*/

/*
$ curl http://127.0.0.1:9080/apisix/admin/services/201  -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -i -d '
{
    "plugins": {
        "limit-count": {
            "count": 2,
            "time_window": 60,
            "rejected_code": 503,
            "key": "remote_addr"
        }
    },
    "enable_websocket": true,
    "upstream": {
        "type": "roundrobin",
        "nodes": {
            "39.97.63.215:80": 1
        }
    }
}'
*/

func (s *Svc) inet() string {

	inetip := s.InetIp

	if inetip == "" {
		inetip = net.GetIntranetIp()[0]
		s.InetIp = inetip
	}

	inetip += ":" + s.Port

	return inetip
}

func (s *Svc) routerExists() bool {

	resp, err := s.Get(ROUTER_URI)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode != 404
}

func (s *Svc) upstreamExists() bool {

	resp, err := s.Get(UPSTREAM_URI)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode != 404
}

func (s *Svc) registerUpstream() (err error) {

	var handler func(string, io.Reader) (*http.Response, error)

	if s.upstreamExists() { // update
		handler = s.Patch
	} else { // insert
		handler = s.Put
	}

	inetip := s.inet()

	upstream := Upstream{
		Name:    s.Name + s.Version,
		Type:    "roundrobin",
		Nodes:   map[string]int{inetip: 1}, // 权重做自增变量
		Retries: 1,
		Checks: &Checks{
			Active: &Active{
				Timeout:  1,
				HTTPPath: "/" + s.Name + "/healthCheck",
				Host:     s.Hosts[0],
				Healthy: &Healthy{
					Interval:     3,
					HTTPStatuses: []int{200},
					Successes:    2,
				},
				Unhealthy: &Unhealthy{
					HTTPStatuses: []int{502, 503, 504},
					HTTPFailures: 2,
					TCPFailures:  3,
					Interval:     3,
				},
				ReqHeaders: []string{"User-Agent: curl/7.29.0"},
			},
			Passive: &Passive{
				Healthy: &Healthy{
					HTTPStatuses: []int{200},
					Successes:    3,
				},
				Unhealthy: &Unhealthy{
					HTTPStatuses: []int{502, 503, 504},
					HTTPFailures: 2,
					TCPFailures:  3,
				},
			},
		},
	}

	if s.HTTP2 {
		upstream.Scheme = "grpc"
	} else {
		upstream.Scheme = "http"
	}

	resp, err := handler(UPSTREAM_URI, encode(upstream))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp)

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	return
}

func (s *Svc) registerService() (err error) {

	var handler func(string, io.Reader) (*http.Response, error)

	if s.routerExists() { // update
		handler = s.Patch
	} else { // insert
		handler = s.Put
	}

	if s.Plugins == nil {
		s.Plugins = map[string]interface{}{}
	}

	service := &Service{
		Name:             s.Name + "+" + s.Version,
		Plugins:          s.Plugins,
		Enable_Websocket: s.EnableWebsocket,
		UpstreamId:       s.Name + "+" + s.Version,
	}

	resp, err := handler(SERVICE_URI, encode(service))

	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Println("Registe service [", s.Name, "] status code:", resp.StatusCode)

	fmt.Println(resp)

	return
}

func (s *Svc) registerRouter(routes string) (err error) {

	var handler func(string, io.Reader) (*http.Response, error)

	// 无论路由是否存在，都按照 put 执行
	// if s.routerExists() { // update
	// handler = s.Patch
	// } else { // insert
	handler = s.Put
	// }

	if s.Hosts == nil {
		s.Hosts = []string{"127.0.0.1"}
	}

	if s.Remote_Addrs == nil {
		s.Remote_Addrs = []string{"0.0.0.0/0"}
	}

	if s.Methods == nil {
		s.Methods = []string{"PUT", "GET", "POST", "PATCH", "DELETE", "OPTIONS", "HEAD", "CONNECT", "TRACE"}
	}

	if s.Plugins == nil {
		s.Plugins = map[string]interface{}{}
	}

	router := &Router{
		Uri:              routes,
		Hosts:            s.Hosts,
		Remote_Addrs:     s.Remote_Addrs,
		Methods:          s.Methods,
		Enable_Websocket: s.EnableWebsocket,
		UpstreamId:       s.Name + s.Version,
		// Service_Id:       s.Name,
		// 直接绑定 upstreamid
		Name: s.Name,
		// Version:          s.Version,
	}

	fmt.Println("go handler", ROUTER_URI)

	resp, err := handler(ROUTER_URI, encode(router))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Registe router [", router, "] status code:", resp.StatusCode)

	fmt.Println(resp)

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	return
}

func encode(d interface{}) io.Reader {
	b, _ := json.Marshal(d)

	return bytes.NewBuffer(b)
}

func decode(resp io.Reader, d interface{}) (err error) {
	b, err := ioutil.ReadAll(resp)
	if err != nil {
		return
	}
	return json.Unmarshal(b, &d)
}
