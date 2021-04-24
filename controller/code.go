package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeTokenInvalid
	CodeRefreshTokenFail

	CodeVoteExpire
	CodeNotRepeatVote
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "参数错误",
	CodeUserExist:        "用户已存在",
	CodeUserNotExist:     "用户不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeServerBusy:       "服务繁忙",
	CodeTokenInvalid:     "token无效",
	CodeRefreshTokenFail: "token刷新失败",
	CodeVoteExpire:       "投票时间已过",
	CodeNotRepeatVote:    "请勿重复投票",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
