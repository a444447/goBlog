package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块的出错 1000开头

	ErrorUsernameUsed     = 1001
	ErrorPasswordWrong    = 1002
	ErrorUsernameNotExist = 1003
	ErrorTokenNotExist    = 1004
	ErrorTokenRuntime     = 1005
	ErrorTokenWrong       = 1006
	ErrorTokenType        = 1007
	ErrorUserNoRight      = 1008
	// 分类模块出错 2000开头
	ErrorCategoryUsed       = 2001
	ErrorCategoryNotExisted = 2002
	// 文章模块出错 3000开头
	ErrorArticleNotExisted = 3001
)

var codeMsg = map[int]string{

	SUCCESS:                 "OK",
	ERROR:                   "FAIL",
	ErrorUsernameUsed:       "用户名已经存在",
	ErrorPasswordWrong:      "密码错误",
	ErrorUsernameNotExist:   "用户不存在",
	ErrorTokenNotExist:      "TOKEN不存在",
	ErrorTokenRuntime:       "TOKEN过期",
	ErrorTokenWrong:         "TOKEN错误",
	ErrorTokenType:          "TOKEN格式错误",
	ErrorUserNoRight:        "用户无权限",
	ErrorCategoryUsed:       "分类已经存在",
	ErrorCategoryNotExisted: "分类不存在",
	ErrorArticleNotExisted:  "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
