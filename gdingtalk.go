package gdingtalk

import (
	"github.com/zjlsliupei/gdingtalk/request"
)

// 实例化Client
func NewClient() request.Client {
	return request.NewClient()
}

// NewRequest
func NewRequest() request.Request {
	return request.NewRequest()
}
