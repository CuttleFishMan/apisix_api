package apisix

type Router struct {
	Uri              string   `json:"uri"`
	Hosts            []string `json:"hosts"`
	Remote_Addrs     []string `json:"remote_addrs"`
	Methods          []string `json:"methods"`
	Enable_Websocket bool     `json:"enable_websocket"`
	UpstreamId       string   `json:"upstream_id,omitempty"`
	Name             string   `json:"name"`
}

type Service struct {
	Plugins          interface{} `json:"plugins"`
	Enable_Websocket bool        `json:"enable_websocket"`
	Name             string      `json:"name"`
	UpstreamId       string      `json:"upstream_id"`
}

type Upstream struct {
	Type       string         `json:"type,omitempty"`
	Nodes      map[string]int `json:"nodes,omitempty"`
	Checks     *Checks        `json:"checks,omitempty"`
	Scheme     string         `json:"scheme,omitempty"`
	Name       string         `json:"name,omitempty"`
	Retries    int            `json:"retries,omitempty"`
	RetryTimes int            `json:"retry_timeout,omitempty"`
}

type Checks struct {
	Passive *Passive `json:"passive"`
	Active  *Active  `json:"active,omitempty"`
}

type Passive struct {
	Healthy   *Healthy   `json:"healthy"`
	Unhealthy *Unhealthy `json:"unhealthy"`
}

type Active struct {
	Timeout    int    `json:"timeout"`
	HTTPPath   string `json:"http_path"`
	Host       string `json:"host"`
	Healthy    *Healthy
	Unhealthy  *Unhealthy
	ReqHeaders []string `json:"req_headers"`
}

type Healthy struct {
	HTTPStatuses []int `json:"http_statuses"`
	Successes    int   `json:"successes"`
	Interval     int   `json:"interval"`
}

type Unhealthy struct {
	Interval     int   `json:"interval"`
	HTTPStatuses []int `json:"http_statuses"`
	HTTPFailures int   `json:"http_failures"`
	TCPFailures  int   `json:"tcp_failures"`
}
