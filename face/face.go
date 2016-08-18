package face

import (
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/execution"
)

type Face interface {
	Init(execution.Execution) error
	GetVersion() (string, error)
	SetLogger(lorg.Logger)
}
