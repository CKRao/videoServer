package defs

//定义Err结构体
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

//定义错误响应结构体
type ErrorResponse struct {
	HttpSC int
	Error  Err
}

//初始化错误响应结构体
var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{
		Error:     "Request body is not correct",
		ErrorCode: "001",
	}}

	ErrorNotAuthUser = ErrorResponse{HttpSC: 401, Error: Err{
		Error:     "User authentication failed",
		ErrorCode: "002",
	}}
)
