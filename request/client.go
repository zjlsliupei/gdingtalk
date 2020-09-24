package request

import "github.com/astaxie/beego/httplib"

const (
	OapiHost = "https://oapi.dingtalk.com"
)

type Client struct {
}

func NewClient() Client {
	return Client{}
}

func (c *Client) Execute(req Request) Response {
	res := c.sendRequest(req)
	return res
}

// sendRequest 发送http请求
func (c *Client) sendRequest(req Request) Response {
	var url string
	queryString := req.queryParams.Encode()
	if queryString == "" {
		url = OapiHost + req.path
	} else {
		url = OapiHost + req.path + "?" + req.queryParams.Encode()
	}

	var httpReq *httplib.BeegoHTTPRequest
	var err error
	if req.method == "post" {
		httpReq = httplib.Post(url)
	} else if req.method == "get" {
		httpReq = httplib.Get(url)
		httpReq.JSONBody(req.bodyParam)
		httpReq.Header("Content-Type", "application/json")
	}
	res, err := httpReq.Response()
	return NewResponse(*res, err)
}
