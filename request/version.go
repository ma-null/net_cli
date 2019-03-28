package request

import (
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
)

func VersionRequest(srv Server) (handlers.VerResponse, error) {
	resp, err := srv.Get( "version")
	if err != nil {
		return handlers.VerResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.VerResponse{}, err
	}

	var apiVer handlers.VerResponse
	if err = json.Unmarshal(body, &apiVer); err != nil {
		return handlers.VerResponse{}, err
	}
	return apiVer, nil
}
