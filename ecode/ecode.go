package ecode

import (
	"gotil/ecode"
)

var (
	Add = ecode.Add
	// service error
	InvalidAppKeyErr = Add(10000) // appkey 验证失败
	RegisterPhoneErr = Add(10001) // 账号格式不对
	UserNameExistErr = Add(10002) // 账号已存在，不能重复注册
)
