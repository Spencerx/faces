package face

import (
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/executor"
)

type Face interface {
	Init(executor.Executor) error
	GetVersion() (string, error)
	SetLogger(lorg.Logger)
}
