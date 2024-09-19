package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("errorCode: %d, msg: %s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(agrs []interface{}) string {
	return fmt.Sprintf(e.msg, agrs...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	NewError := *e
	NewError.details = []string{}
	/*for _, d := range details {
		NewError.details = append(NewError.details, d)
	}*/
	NewError.details = append(NewError.details, details...)
	fmt.Println(NewError, len(NewError.details))
	return &NewError
}

func (e *Error) StateCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
