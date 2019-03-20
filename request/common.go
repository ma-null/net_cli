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

func CreateRequestURL(pr Params, requestPath string) string {
	if pr.NetIfVersion != "" {
		return strings.Join([]string{"http://" + pr.Server + ":" + pr.Port, "service", pr.NetIfVersion, requestPath}, "/")
	}
	return strings.Join([]string{"http://" + pr.Server + ":" + pr.Port, "service", requestPath}, "/")
}

func MakeRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
