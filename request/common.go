package request

import (
	"net/http"
	"strings"
)

type Params struct {
	NetIfVersion string
	Server       string
	Port         string
}

func createRequestURL(pr Params, requestPath string) string {
	if pr.NetIfVersion != "" {
		return strings.Join([]string{"http://" + pr.Server + ":" + pr.Port, "service", pr.NetIfVersion, requestPath}, "/")
	}
	return strings.Join([]string{"http://" + pr.Server + ":" + pr.Port, "service", requestPath}, "/")
}

type (
	Server interface {
		Get(requestPath string) (*http.Response, error)
		UpdateVersion() error
	}

	server struct {
		pr Params
	}
)

func NewServer(pr Params) Server {
	return &server{pr}
}

func (s *server) Get(requestPath string) (*http.Response, error) {
	reqURL := createRequestURL(s.pr, requestPath)
	return http.Get(reqURL)
}

func (s *server) UpdateVersion() error {
	ver, err := VersionRequest(s)
	if err != nil {
		return err
	}
	s.pr.NetIfVersion = ver.Version
	return nil
}

