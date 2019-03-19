package request

import (
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
)

func InterfaceListRequest(pr Params) (handlers.IfResponse, error) {
	resp, err := MakeRequest(CreateRequestURL(pr, "interfaces"))
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
