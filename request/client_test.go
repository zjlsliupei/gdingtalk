package request

import (
	"log"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	c := NewClient()
	req := NewRequest()
	req.SetPath("/gettoken")
	req.SetMethod("Get")
	param := url.Values{}
	param.Add("appkey", "dingwlgogldfur1gw2ma")
	param.Add("appsecret", "f6TqiRiOy7QKmyAemBmlUjOtbSe2QY0BuwGb_P-9Lf9lt1OiokQf9uuUGan2H80Q")
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
