package timeout

import (
	"net/http"
	"time"
)

type CallBackFunc func(*http.Request)
type MsgCallBackFunc func(*http.Request) interface{}

type Option func(*TimeoutWriter)

type TimeoutOptions struct {
	CallBack      CallBackFunc
	DefaultMsg    interface{}
	MsgCallBack   MsgCallBackFunc
	Timeout       time.Duration
	ErrorHttpCode int
}

func WithTimeout(d time.Duration) Option {
	return func(t *TimeoutWriter) {
		t.Timeout = d
	}
}

// Optional parameters
func WithErrorHttpCode(code int) Option {
	return func(t *TimeoutWriter) {
		t.ErrorHttpCode = code
	}
}

// Optional parameters
func WithDefaultMsg(resp interface{}) Option {
	return func(t *TimeoutWriter) {
		t.DefaultMsg = resp
	}
}

// Optional parameters
func WithCallBack(f CallBackFunc) Option {
	return func(t *TimeoutWriter) {
		t.CallBack = f
	}
}

// Optional parameters
func WithMsgCallBack(f MsgCallBackFunc) Option {
	return func(t *TimeoutWriter) {
		t.MsgCallBack = f
	}
}
