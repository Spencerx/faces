package faces

import (
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/logger"
)

func SetLogger(log lorg.Logger) {
	logger.Logger = log
}
