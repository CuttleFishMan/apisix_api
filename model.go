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
}
