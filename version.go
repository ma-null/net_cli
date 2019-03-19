package main

import (
  
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/ma-null/NetInterface/handlers"

)



func MakeVersionRequest() (handlers.VerResponse, error) {
	resp, err := http.Get("http://localhost:8080/service/version")
	if err != nil {
		return  handlers.VerResponse{}, err
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.VerResponse{}, err
  }
  var apiVer handlers.VerResponse
  err = json.Unmarshal(body, &apiVer)
	if err != nil {
		return handlers.VerResponse{}, err
	}
  
	return apiVer, nil
}
