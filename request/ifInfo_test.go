package request

import (
	"bytes"
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
	"net/http"
	"testing"
)

var netInterfaceInfo = handlers.NetInterface{
	"eth0",
	[]byte {2},
	[]string{"172.17.128.27/24"},
	1500,
}

func mockGetInfo(_ string) (*http.Response, error) {

	nI,_:= json.Marshal(netInterfaceInfo)
	return &http.Response{
		Status:	"200 OK",
		StatusCode: 200,
		Body: ioutil.NopCloser(bytes.NewReader(nI)),
	}, nil

}

func TestInterfaceInfoRequest(t *testing.T) {
	pr := MockParams()
	srv := NewMockServer(pr, mockGetInfo)
	resp, err := InterfaceInfoRequest("", srv)
	if  err != nil {
		t.Error(err.Error())
	}
	if resp.Name != "eth0" {
		t.Error("Expected: " + netInterfaceInfo.Name + ", got: "+ resp.Name)
	}
}