package e

var code2Msg = map[int]string{
	SUCCESS:     "ok",
	INNER_ERROR: "内部错误",
	BAD_REQUEST: "请求有问题",
}

func GetMsg(code int) string {
	return code2Msg[code]
}
