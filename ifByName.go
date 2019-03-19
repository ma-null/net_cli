package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/ma-null/NetInterface/handlers"

)



func MakeInterfaceRequest(name string, pr Params) (handlers.NetInterface, error) {
  
	resp, err := http.Get("http://"+pr.Server+":"+pr.Port+"/service/"+pr.NetIfVersion+"/interface/"+name) // to do server and port is a key
	if err != nil {
    fmt.Println(err.Error())
		return handlers.NetInterface{}, err
	}

  
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.NetInterface{}, err
	}

  var netIf handlers.NetInterface
  err = json.Unmarshal(body, &netIf)
	if err != nil {
		return handlers.NetInterface{}, err
	}
  
	return netIf, nil
	
}
