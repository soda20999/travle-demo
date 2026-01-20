package response

type ResCode int64

const (
	CodeSuccess            ResCode = 1000 + iota //成功
	CodeInvalidParams                            //参数错误
	CodeUserExist                                //用户已存在
	CodeUserNotExist                             //用户不存在
	CodeInvalidPassword                          //密码错误
	CodeServerBusy                               //服务繁忙
	CodeTokenHeaderEmpty                         //请求头中Token为空
	CodeTokenInvalidFormat                       //请求头中Token格式有误
	CodeTokenInvalid                             //无效的Token
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:            "成功",
	CodeInvalidParams:      "参数错误",
	CodeUserExist:          "用户已存在",
	CodeUserNotExist:       "用户不存在",
	CodeInvalidPassword:    "密码错误",
	CodeServerBusy:         "服务繁忙",
	CodeTokenHeaderEmpty:   "请求头中Token为空",
	CodeTokenInvalidFormat: "请求头中Token格式有误",
	CodeTokenInvalid:       "无效的Token",
}

func getMsg(code ResCode) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
