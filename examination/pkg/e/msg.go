package e
//此包根据错误码返回信息
var MsgFlags = map[int]string {
SUCCESS : "ok",
ERROR : "fail",
INVALID_PARAMS :"请求参数错误",
ERROR_POWER:"用户权限不足",
ERROR_CODE:"验证码错误",
ERROR_COMMENT: "文章不可评论",
ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
ERROR_AUTH_CHECK_TOKEN_FAIL : "Token解析失败",
ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
ERROR_AUTH_TOKEN : "Token生成失败",
ERROR_AUTH : "账号或密码错误",
ERROR_REGISTER :"账号已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
