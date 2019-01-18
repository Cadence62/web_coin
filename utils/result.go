package utils

import "time"

// 返回值结构定义
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Time int64       `json:"time"`
	Data interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Time: time.Now().Unix(),
		Data: nil,
	}
}
