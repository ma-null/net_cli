package request

type Params struct {
  NetIfVersion string
  Server string
  Port  string

}

func CreateRequest(pr Params) string {
	return "http://"+pr.Server+":"+pr.Port+"/service/"+pr.NetIfVersion+"/interfaces"
}