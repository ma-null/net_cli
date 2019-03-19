package request

import (
	"net/http"
)

type Params struct {
  NetIfVersion string
  Server string
  Port  string
}

func CreateRequest(pr Params, str string) (*http.Response, error) {
	url:=""
	if str =="/version"{
		url="http://localhost:8080/service/"+str	
	} else {
		url="http://"+pr.Server+":"+pr.Port+"/service/"+str	
	}	
	resp, err := http.Get(url)
	if err != nil {
		return  nil, err
	}
	return resp, nil
	
}