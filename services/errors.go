package services

import "github.com/astaxie/beego/logs"

func ErrorCheck(msg string, err error) {
	if err != nil {
		logs.Error(msg, " Reason: ", err)
	}
}

func ErrorCheckWithSuccess(errMsg, passMsg string, err error) {
	if err != nil {
		logs.Debug(errMsg, " Reason: ", err)
	} else {
		logs.Debug(passMsg)
	}
}
