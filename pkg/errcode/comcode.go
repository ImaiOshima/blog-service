package errcode

var (
	Success                     = NewError(0, "成功")
	ServerError                 = NewError(10000000, "服务内部错误")
	InvalidParam                = NewError(10000001, "入参错误")
	NotFound                    = NewError(10000002, "找不到")
	UnauthoronizedAuthNotExist  = NewError(10000003, "鉴权失败")
	UnauthoronizedTokenError    = NewError(10000004, "鉴权失败,Token失败")
	UnauthoronizedTokenTimeout  = NewError(10000005, "鉴权失败,Token超时")
	UnauthoronizedTokenGenerate = NewError(10000006, "鉴权失败,Token生成失败")
	TooManyRequests             = NewError(10000007, "请求过")
)
