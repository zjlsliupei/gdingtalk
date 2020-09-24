package request

import (
	"errors"
	"net/url"
	"strings"
)

type Request struct {
	method      string
	path        string
	queryParams url.Values
	bodyParam   map[string]interface{}
	contentType string
}

// NewRequest
func NewRequest() Request {
	return Request{
		contentType: "application/json",
	}
}

// SetMethod 设置请求类型，支持post,get
// sample: SetMethod("POST") 或 SetMethod("post") 都是合法
func (req *Request) SetMethod(method string) error {
	lowerMethod := strings.ToLower(method)
	if (lowerMethod != "post") && (lowerMethod != "get") {
		return errors.New("method: " + method + " is invalid")
	}
	req.method = lowerMethod
	return nil
}

// SetPath 设置请求路
// sample:  SetPath("/get_jsapi_ticket")
func (req *Request) SetPath(path string) error {
	req.path = path
	return nil
}

// SetQueryParam 设置query参数
func (req *Request) SetQueryParam(values url.Values) error {
	req.queryParams = values
	return nil
}

// SetBodyParam 设置post body参数
// sample:
// SetBodyParam(map[string]interface{}{
//     "name": "sven",
//     "age": 18
// })
func (req *Request) SetBodyParam(body map[string]interface{}) error {
	req.bodyParam = body
	return nil
}
