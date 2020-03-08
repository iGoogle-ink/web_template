package ecode

import (
	"strconv"
	"sync"
)

var (
	errorMap     = new(sync.Map)
	OK           = add(0, "SUCCESS")
	CookieErr    = add(302, "Cookie 失效")
	RequestErr   = add(400, "请求错误")
	Unauthorized = add(401, "未授权")
	NothingFound = add(404, "啥都木有")
	ServerErr    = add(500, "服务器错误")
	ServerBusy   = add(502, "服务器忙，请稍微尝试")

	// todo some error code and msg to add
)

// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
}

func add(code int, msg string) Error {
	errorMap.Store(code, msg)
	return Error(code)
}

func New(code int, msg string) Error {
	errorMap.Store(code, msg)
	return Error(code)
}

// A Code is an int error code spec.
type Error int

func (e Error) Error() string { return strconv.Itoa(int(e)) }

// Code return error code
func (e Error) Code() int { return int(e) }

// Message return error message
func (e Error) Message() string {
	if msg, ok := errorMap.Load(e.Code()); ok {
		return msg.(string)
	}
	return e.Error()
}

// analyse error info
func AnalyseError(err error) Codes {
	if err == nil {
		return OK
	}
	if codes, ok := err.(Codes); ok {
		return codes
	}
	return errStringToError(err.Error())
}

func errStringToError(e string) Error {
	if e == "" {
		return OK
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		return add(-1, e)
	}
	return Error(i)
}
