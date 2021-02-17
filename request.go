package apisix

import (
	"io"
	"net/http"
)

func (s *Svc) Get(uri string) (resp *http.Response, err error) {
	return s.request(http.MethodGet, uri, nil)
}

func (s *Svc) Put(uri string, body io.Reader) (resp *http.Response, err error) {
	return s.request(http.MethodPut, uri, body)
}

func (s *Svc) Patch(uri string, body io.Reader) (resp *http.Response, err error) {
	return s.request(http.MethodPatch, uri, body)
}

func (s *Svc) request(method, uri string, body io.Reader) (resp *http.Response, err error) {

	apisixHost := s.Host
	if apisixHost == "" {
		apisixHost = "127.0.0.1:9180"
	}

	req, err := http.NewRequest(method, "http://"+apisixHost+uri+s.Version, body)

	if err != nil {
		return
	}
	req.Header.Add("X-API-KEY", s.XAPIKEY)
	c := &http.Client{}

	return c.Do(req)
}
