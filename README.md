# gdingtalk
dingtalk for glang

## 安装
```go
go get github.com/zjlsliupei/gdingtalk
```

## 快速开始
```go
import (
 "github.com/zjlsliupei/gdingtalk"
)


c := gdingtalk.NewClient()
// 构造请求参数
req := gdingtalk.NewRequest()
req.SetPath("/get_jsapi_ticket")
req.SetMethod("Get")
// 执行请求
res := c.Execute(req)
res.IsSuccess() // false
res.GetBodyData("errcode").Int() // 100000
res.GetError()) // 不合法的access_token
```

## 通讯录模块
### 获取企业下所有用户
```go
import (
 "github.com/zjlsliupei/gdingtalk"
)

u := gdingtalk.NewUser()
users, err := u.GetAllUsers()
```