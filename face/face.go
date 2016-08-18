package face

import (
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/execution"
)

type Interface interface {
	Init(execution.Execution) error
	GetVersion() (string, error)
	SetLogger(lorg.Logger)
}

type Abstract struct {
	Logger lorg.Logger
}

func (abstract *Abstract) SetLogger(log lorg.Logger) {
	abstract.Logger = log
}
