package request

import (
  "io/ioutil"
  "encoding/json"
  "github.com/ma-null/NetInterface/handlers"
)

func InterfaceInfoRequest(name string, pr Params) (handlers.NetInterface, error) {
	str:= pr.NetIfVersion+"/interface/"+name
	resp, err := CreateRequest(pr, str)
	if err != nil {
		return handlers.NetInterface{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handlers.NetInterface{}, err
	}

  var netIf handlers.NetInterface
	if err = json.Unmarshal(body, &netIf); err != nil {
		return handlers.NetInterface{}, err
	}
	return netIf, nil
}
