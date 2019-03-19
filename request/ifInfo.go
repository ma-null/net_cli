package request

import (
	"encoding/json"
	"github.com/ma-null/NetInterface/handlers"
	"io/ioutil"
)

func InterfaceInfoRequest(name string, pr Params) (handlers.NetInterface, error) {
	resp, err := MakeRequest(CreateRequestURL(pr, "interface/" + name))
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
