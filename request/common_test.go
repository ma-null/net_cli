package request

import (
	"net/http"
)

func MockParams() Params {
	return Params{
		"v1",
		"127.0.0.1",
		"8080",
	}
}

type mockServer struct {
	pr Params
	mockGet func(requestPath string) (*http.Response, error)
}

func NewMockServer(pr Params, mockGet func(requestPath string) (*http.Response, error)) Server {
	return &mockServer{pr, mockGet}
}

func (s *mockServer) Get(requestPath string) (*http.Response, error) {
	return s.mockGet(requestPath)
}

func (s *mockServer) UpdateVersion() error {
	s.pr.NetIfVersion = "v1"
	return nil
}
