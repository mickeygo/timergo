package infrastructure

import "errors"

var (
	// Method_Unimplement_Error 函数未实现错误
	Method_Unimplement_Error error = errors.New("Method unimplement.")
)