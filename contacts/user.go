package contacts

import (
	"github.com/tidwall/gjson"
	"github.com/zjlsliupei/gdingtalk/request"
	"net/url"
	"strconv"
)

type User struct {
	AccessToken string
}

func (u *User) GetAllUsers() []gjson.Result {
	// deptId=1,为全公司
	return u.GetUsersByDeptId(1)
}

// GetUsersByDeptId 获取部门下用户
func (u *User) GetUsersByDeptId(deptId int64) []gjson.Result {
	department := Department{AccessToken: u.AccessToken}
	deps, err := department.GetDepartments(deptId)
	var users []gjson.Result
	if err != nil {
		return users
	}
	for _, v := range deps {
		userIds := u.GetUserIdByDeptId(v.Get("id").Int())
		if len(userIds) > 0 {
			for _, v := range userIds {
				user := u.GetUserDetail(v)
				if user.Exists() {
					users = append(users, user)
				}
			}
		}
	}
	// 加入根据目录下所属用户
	userIds := u.GetUserIdByDeptId(deptId)
	if len(userIds) > 0 {
		for _, v := range userIds {
			user := u.GetUserDetail(v)
			if user.Exists() {
				users = append(users, user)
			}
		}
	}
	return users
}

// GetUserIdByDeptId 根据部门获取user_ids
func (u *User) GetUserIdByDeptId(deptId int64) []string {
	c := request.NewClient()
	req := request.NewRequest()
	req.SetMethod("GET")
	param := url.Values{}
	param.Add("access_token", u.AccessToken)
	param.Add("deptId", strconv.FormatInt(deptId, 10))
	req.SetQueryParam(param)
	req.SetPath("/user/getDeptMember")
	res := c.Execute(req)
	var userIds []string
	if !res.IsSuccess() {
		return userIds
	}
	res.GetBodyData("userIds").ForEach(func(key, value gjson.Result) bool {
		userIds = append(userIds, value.String())
		return true
	})
	return userIds
}

// GetUserDetail 根据user_id获取用户详情
func (u *User) GetUserDetail(userId string) gjson.Result {
	c := request.NewClient()
	req := request.NewRequest()
	req.SetMethod("GET")
	param := url.Values{}
	param.Add("access_token", u.AccessToken)
	param.Add("userid", userId)
	req.SetQueryParam(param)
	req.SetPath("/user/get")
	res := c.Execute(req)
	if !res.IsSuccess() {
		return gjson.Result{}
	}
	return res.GetBodyData()
}
