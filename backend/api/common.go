package api

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var (
	// common errors
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrUnauthorizedAP      = newError(401, "账号或者密码不正确")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "邮箱地址已被使用")
	ErrRePasswd        = newError(1001, "两次密码输入不正确")
	ErrPhone           = newError(1001, "用户手机号为空或者不正确,请更新手机号")

	ErrCashRequest        = newError(2002, "账户余额不足,提现发起失败")
	ErrCashAccountRequest = newError(2003, "账户状态异常,提现发起失败,请联系客服")
	ErrCashPicRequest     = newError(2004, "账户提现方式为空,提现发起失败,请填写提现方式")
	ErrCashNow            = newError(2005, "有未结束的提现申请,请等待之前申请结束后,重新发起")
)

type Error struct {
	Code    int
	Message string
}
func (err Error) Error() string {
	return err.Message
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}

func HandleSuccess(ctx *gin.Context, message string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	response := Response{
		Code: 0, 
		Message: message, 
		Data: data,
	}

	ctx.JSON(http.StatusOK, response)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data any) {
	if data == nil {
		data = map[string]string{}
	}

	response := Response{
		Code: errorCodeMap[err], 
		Message: err.Error(), 
		Data: data,
	}

	ctx.JSON(httpCode, response)
}

func GetClientIP(request *http.Request) string {
    // Try Forwarded header (RFC 7239 standard)

    if forwarded := request.Header.Get("Forwarded"); forwarded != "" {
        re := regexp.MustCompile(`for="?([^";,]+)`)
        matches := re.FindStringSubmatch(forwarded)
        if len(matches) > 1 {
            return matches[1]
        }
    }

    // Fallback to X-Forwarded-For (common de facto standard)
    if xff := request.Header.Get("X-Forwarded-For"); xff != "" {
        ips := strings.Split(xff, ",")
        if len(ips) > 0 {
            return strings.TrimSpace(ips[0])
        }
    }

    // Final fallback to remote address
    remoteAddr := request.RemoteAddr
    if strings.Contains(remoteAddr, "[") { // IPv6 address
        return strings.Split(remoteAddr, "]")[0] + "]"
    }
    return strings.Split(remoteAddr, ":")[0] // IPv4 address
}
