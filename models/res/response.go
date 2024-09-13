package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	Success = 0
	Failed  = 1
)

func Result(code int, message string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Ok 请求成功
func Ok(data any, message string, c *gin.Context) {
	Result(Success, message, data, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, "操作成功", data, c)
}

func OkWithMessage(data any, message string, c *gin.Context) {
	Result(Success, message, map[string]any{}, c)
}

func OkWith(c *gin.Context) {
	Result(Success, "成功~", map[string]any{}, c)
}

// Fail 请求失败
func Fail(data any, message string, c *gin.Context) {
	Result(Failed, message, data, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Failed, message, map[string]any{}, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), msg, map[string]any{}, c)
		return
	}
	Result(Failed, msg, map[string]any{}, c)
}
