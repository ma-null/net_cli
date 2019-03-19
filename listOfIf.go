package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/ma-null/NetInterface/handlers"

)

func MakeListRequest(pr Params) (handlers.IfResponse, error) {
	resp, err := http.Get("http://"+pr.Server+":"+pr.Port+"/service/"+pr.NetIfVersion+"/interfaces")
	if err != nil {
		return  handlers.IfResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.IfResponse{}, err
  }

  var listOfNames handlers.IfResponse
  err = json.Unmarshal(body, &listOfNames)
	if err != nil {
		return handlers.IfResponse{}, err
	}
  
	return listOfNames, nil
}
