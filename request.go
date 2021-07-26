package apisix

import (
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
		return s.handleVersion(uri)
	}

	return s.handleVersion(s.Prefix + uri)
}

func (s *Svc) handleVersion(uri string) string {

	if s.Version != "" {
		return uri + "/" + s.Version
	}

	return uri
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
