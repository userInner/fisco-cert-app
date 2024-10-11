package controller

type ResCode uint64

const (
	CtxUserIDKey          = "userID"
	CtxUserAccessTokenKey = "userAToken"
	CtxUserAddrKey        = "userAddr"
)

const (
	CodeSuccess ResCode = iota + 1000
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidUploadFile
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码不存在",
	CodeServerBusy:        "服务繁忙",
	CodeNeedLogin:         "需要登陆",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidUploadFile: "上传失败",
}

func (r ResCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
