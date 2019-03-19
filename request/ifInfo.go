package request

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/ma-null/NetInterface/handlers"
)

func InterfaceInfoRequest(name string, pr Params) (handlers.NetInterface, error) {
	resp, err := http.Get(CreateRequest(pr)) // to do server and port is a key
	if err != nil {
    fmt.Println(err.Error())
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
