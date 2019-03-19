package request

import (
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/ma-null/NetInterface/handlers"
)

func VersionRequest(pr Params) (handlers.VerResponse, error) {
	resp, err := http.Get(CreateRequest(pr))
	if err != nil {
		return  handlers.VerResponse{}, err
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
