package v1

import (
	"errors"
	"net/http"
)

var (
	// ErrorValidation 表单验证错误
	ErrorValidation = errors.New("表单验证错误")
	// ErrorUpdate 更新失败
	ErrorUpdate = errors.New("更新路由数据失败")
)

var (
	// ErrorCode 错误
	ErrorCode = http.StatusNotFound
	// SuccessCode 成功
	SuccessCode = http.StatusOK
)

// GetDataSuccess 获取成功
var GetDataSuccess = "获取成功"

// UpdateDataSuccess 更新成功
var UpdateDataSuccess = "更新成功"
