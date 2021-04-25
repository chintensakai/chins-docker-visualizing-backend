package pkg

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "error",
	INVALID_PARAMS:              "参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
	ERROR_AUTH_TOKEN:            "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
