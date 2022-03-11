package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回自定义状态码
const (
	Successful       = 0
	PermissionDenied = 403
	NotFound         = 404
	ManyRequest      = 429
	Failed           = 500
	AuthFail         = 401
)

// Response HTTP返回数据结构体, 可使用这个, 也可以自定义
type Response struct {
	Code    int         `json:"status"` // 状态码,这个状态码是与前端和APP约定的状态码,非HTTP状态码
	Data    interface{} `json:"data"`   // 返回数据
	Message string      `json:"msg"`    // 自定义返回的消息内容
	msgData interface{} // 消息解析使用的数据
}

// Success 获得一个基本的表示请求成功的 Response 对象
func Success(data interface{}) *Response {
	return &Response{Code: Successful, Data: data, Message: "success"}
}

// Fail 获得一个基本的表示请求失败的 Response 对象
func Fail(data interface{}) *Response {
	return &Response{Code: Failed, Data: data, Message: "fail"}
}

func Err(err error) *Response {
	return &Response{Code: Failed, Data: nil, Message: err.Error()}
}

// WithCode 获得一个 Response 对象
func WithCode(code int, data interface{}) *Response {
	return &Response{Code: code, Data: data, Message: "success"}
}

// Msg msg 描述
func (rsp *Response) Msg(msg string) *Response {
	rsp.Message = msg
	return rsp
}

// MsgData 消息解析使用的数据
func (rsp *Response) MsgData(data interface{}) *Response {
	rsp.msgData = data
	return rsp
}

// End 在调用了这个方法之后,还是需要 return 的
func (rsp *Response) End(c *gin.Context, httpStatus ...int) {
	status := http.StatusOK
	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}

	c.JSON(status, rsp)
}

// Object 直接获得本对象
func (rsp *Response) Object(_ *gin.Context) *Response {
	return rsp
}

func NewResponse(code int, data interface{}, message ...string) *Response {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	return &Response{Code: code, Data: data, Message: msg}
}
