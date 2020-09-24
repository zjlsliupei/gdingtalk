package request

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	httpCode     int
	httpResponse http.Response
	httpErr      string
	errCode      int64
	errMsg       string
	httpBody     string
	success      bool
}

// NewResponse 实例化response
func NewResponse(httpResponse http.Response, reqErr error) Response {
	res := Response{
		httpResponse: httpResponse,
	}
	if reqErr != nil {
		res.httpErr = reqErr.Error()
	}
	res.init()
	return res
}

// init 解析钉钉返回
func (res *Response) init() {
	body, err := ioutil.ReadAll(res.httpResponse.Body)
	defer res.httpResponse.Body.Close()
	if err != nil {
		res.success = false
		res.httpErr = err.Error()
		log.Println("dingtalk http request error:", res.httpResponse.Request.URL.String(), res.httpErr)
		return
	}
	res.httpBody = string(body)
	res.errCode = gjson.Get(res.httpBody, "errcode").Int()
	if (res.httpResponse.StatusCode >= 200 && res.httpResponse.StatusCode < 400) && res.errCode == 0 {
		res.success = true
	} else {
		res.success = false
	}
	res.errMsg = gjson.Get(res.httpBody, "errmsg").String()
	if res.success == false {
		log.Println("dingtalk response error:", res.httpResponse.Request.URL.String(), res.errMsg)
	}
}

// IsSuccess 判断返回值是否成功
func (res *Response) IsSuccess() bool {
	return res.success
}

// GetHttpCode 返回httpCode
func (res *Response) GetHttpCode() int {
	return res.httpResponse.StatusCode
}

// GetBodyData 返回钉钉返回数据
func (res *Response) GetBodyData(path ...string) gjson.Result {
	if len(path) == 0 {
		return gjson.Parse(res.httpBody)
	}
	return gjson.Get(res.httpBody, path[0])
}

// GetError 返回错误信息，优先级：httpErr > errMsg
func (res *Response) GetError() string {
	if res.httpErr != "" {
		return res.httpErr
	}
	return res.errMsg
}
