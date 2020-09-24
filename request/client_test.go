package request

import (
	"log"
	"net/url"
	"os"
	"testing"
)

var (
	appKey    string
	appSecret string
)

func init() {
	appKey = os.Getenv("TEST_APPKEY")
	appSecret = os.Getenv("TEST_APPSECRET")
}

func TestGet(t *testing.T) {
	c := NewClient()
	req := NewRequest()
	req.SetPath("/gettoken")
	req.SetMethod("Get")
	param := url.Values{}
	param.Add("appkey", appKey)
	param.Add("appsecret", appSecret)
	req.SetQueryParam(param)
	res := c.Execute(req)
	log.Println(res.IsSuccess(), res.GetBodyData())
}

func TestPost(t *testing.T) {
	c := NewClient()
	req := NewRequest()
	req.SetPath("/user/create")
	query := url.Values{}
	query.Add("access_token", "11111")
	req.SetQueryParam(query)
	req.SetBodyParam(map[string]interface{}{
		"userid": "11111",
	})
	req.SetMethod("Post")
	res := c.Execute(req)
	log.Println(res.IsSuccess(), res.GetError())
}
