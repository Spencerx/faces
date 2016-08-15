package faces

import "github.com/kovetskiy/lorg"

var logger = lorg.NewDiscarder()

func debugf(format string, values ...interface{}) {
	logger.Debugf(format, values...)
}

func debugln(value interface{}) {
	logger.Debug(value)
}
