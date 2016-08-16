package logger

import "github.com/kovetskiy/lorg"

var Logger = lorg.NewDiscarder()

func Debugf(format string, values ...interface{}) {
	Logger.Debugf(format, values...)
}

func Tracef(format string, values ...interface{}) {
	Logger.Tracef(format, values...)
}
