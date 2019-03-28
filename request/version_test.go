package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func mockGetVersion(_ string) (*http.Response, error) {
	return &http.Response{
		Status:	"200 OK",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewBufferString("{\"version\":\"v1\"}")),
	}, nil
}

func TestVersionRequest(t *testing.T) {
	pr := MockParams()
	srv := NewMockServer(pr, mockGetVersion)
	resp, err := VersionRequest(srv)
	if  err != nil {
		t.Error(err.Error())
	}
	if resp.Version != "v1" {
		t.Error("Expected: {\"version\":\"v1\"}, got: "+ resp.Version)
	}
}
