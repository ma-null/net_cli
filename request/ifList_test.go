package request

import (
	"bytes"
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
	"net/http"
	"testing"
)

var netInterfaceNames = handlers.IfResponse{
	[]string{"eth0","ln0", "smth"},
}

func mockGetNames(_ string) (*http.Response, error) {
	nI,_:= json.Marshal(netInterfaceNames)
	return &http.Response{
		Status:	"200 OK",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewReader(nI)),
	}, nil
}

func TestInterfaceListRequest(t *testing.T) {
	pr := MockParams()
	srv := NewMockServer(pr, mockGetNames)
	resp, err := InterfaceListRequest(srv)
	if  err != nil {
		t.Error(err.Error())
	}
	for i := range resp.AllIntr {
		if resp.AllIntr[i] != netInterfaceNames.AllIntr[i] {
			t.Error("Expected: " + netInterfaceNames.AllIntr[i] + ", got: "+ resp.AllIntr[i])
		}
	}
}