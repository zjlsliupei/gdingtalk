package contacts

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/zjlsliupei/gdingtalk/request"
	"net/url"
	"strconv"
)

type Department struct {
	AccessToken string
}

// GetDepartments 获取部门列表
func (d *Department) GetDepartments(parentId int64) ([]gjson.Result, error) {
	c := request.NewClient()
	req := request.NewRequest()
	req.SetMethod("GET")
	param := url.Values{}
	param.Add("access_token", d.AccessToken)
	param.Add("fetch_child", "1")
	param.Add("id", strconv.FormatInt(parentId, 10))
	req.SetQueryParam(param)
	req.SetPath("/department/list")
	res := c.Execute(req)
	if !res.IsSuccess() {
		return nil, errors.New(res.GetError())
	}
	return res.GetBodyData("department").Array(), nil
}
