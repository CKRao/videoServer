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

	ErrorDBError = ErrorResponse{HttpSC: 500, Error: Err{
		Error:     "DB ops failed",
		ErrorCode: "003",
	}}

	ErrorInternalError = ErrorResponse{HttpSC: 500, Error: Err{
		Error:     "Internal server failed",
		ErrorCode: "004",
	}}

	ErrorParamsNullError = ErrorResponse{HttpSC: 401, Error: Err{
		Error:     "Params is Null. Please Check The Params.",
		ErrorCode: "005",
	}}

	ErrorUserNotFoundError = ErrorResponse{HttpSC: 401, Error: Err{
		Error:     "User Not Found",
		ErrorCode: "006",
	}}

	ErrorPasswordWrongError = ErrorResponse{HttpSC: 401, Error: Err{
		Error:     "Password is Wrong",
		ErrorCode: "007",
	}}
)
