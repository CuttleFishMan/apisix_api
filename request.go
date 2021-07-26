package apisix

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Svc) Get(uri string) (resp *http.Response, err error) {
	return s.request(http.MethodGet, s.handlePrefix(uri), nil)
}

func (s *Svc) Put(uri string, body io.Reader) (resp *http.Response, err error) {
	return s.request(http.MethodPut, s.handlePrefix(uri), body)
}

func (s *Svc) Patch(uri string, body io.Reader) (resp *http.Response, err error) {
	return s.request(http.MethodPatch, s.handlePrefix(uri), body)
}

func (s *Svc) handlePrefix(uri string) string {

	if s.Prefix == "" || s.Prefix == "/" {
		// [uri]/upstream+1.3
		return uri + "/" + s.Name + s.Version
	}

	return s.Prefix + uri + "/" + s.Name + s.Version
}

func (s *Svc) request(method, uri string, body io.Reader) (resp *http.Response, err error) {

	apisixHost := s.Host
	if apisixHost == "" {
		apisixHost = "127.0.0.1:9180"
	}

	uri = "http://" + apisixHost + uri
	fmt.Println("request apisix:", method, uri)

	fmt.Println(body)

	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		return
	}
	req.Header.Add("X-API-KEY", s.XAPIKEY)
	c := &http.Client{}

	return c.Do(req)
}
