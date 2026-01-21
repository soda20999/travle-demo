package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 100, //业务状态码
	"message": "ok", //业务提示信息
	"data": {} //业务数据
}
*/

type Response struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseError(code ResCode, c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: getMsg(code),
		Data:    nil,
	})
}

// 参数顺序：context 在前，data 在后
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    CodeSuccess,
		Message: getMsg(CodeSuccess),
		Data:    data,
	})
}

func ResponseErrorWithMsg(message interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code:    CodeSuccess,
		Message: getMsg(CodeSuccess),
		Data:    message,
	})
}
