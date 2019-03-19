package request

import (
  "io/ioutil"
  "encoding/json"
  "github.com/ma-null/NetInterface/handlers"
)

func InterfaceListRequest(pr Params) (handlers.IfResponse, error) {
	str:= pr.NetIfVersion+"/interfaces"
	resp, err := CreateRequest(pr,str)
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
