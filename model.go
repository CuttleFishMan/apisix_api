package apisix

type Router struct {
	Uri              string   `json:"uri"`
	Hosts            []string `json:"hosts"`
	Remote_Addrs     []string `json:"remote_addrs"`
	Methods          []string `json:"methods"`
	Enable_Websocket bool     `json:"enable_websocket"`
	Service_Id       string   `json:"service_id,omitempty"`
}

type Service struct {
	Plugins          interface{} `json:"plugins"`
	Enable_Websocket bool        `json:"enable_websocket"`
	Upstream         Upstream    `json:"upstream"`
}

type Upstream struct {
	Type        string         `json:"type,omitempty"`
	Nodes       map[string]int `json:"nodes,omitempty"`
	ServiceName string         `json:"service_name,omitempty"`
	Passive     *Passive       `json:"passive,omitempty"`
	Scheme      string         `json:"scheme,omitempty"`
}

type Passive struct {
	Healthy   *Healthy   `json:"healthy"`
	Unhealthy *Unhealthy `json:"unhealthy"`
}

type Healthy struct {
	HTTPStatuses []int `json:"http_statuses"`
	Successes    int   `json:"successes"`
}

type Unhealthy struct {
	HTTPStatuses []int `json:"http_statuses"`
	HTTPFailures int   `json:"http_failures"`
	TCPFailures  int   `json:"tcp_failures"`
}
