package request

import (
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
)

func InterfaceListRequest(srv Server) (handlers.IfResponse, error) {
	resp, err := srv.Get( "interfaces")
	if err != nil {
		return handlers.IfResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.IfResponse{}, err
	}

	var listOfNames handlers.IfResponse
	if err = json.Unmarshal(body, &listOfNames); err != nil {
		return handlers.IfResponse{}, err
	}
	return listOfNames, nil
}
