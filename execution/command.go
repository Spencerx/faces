package execution

import (
	"github.com/reconquest/faces/logger"
	"github.com/reconquest/lexec-go"
)

type Execution struct {
	Sudo bool
}

func (execution *Execution) Exec(
	name string, args ...string,
) *Operation {
	if execution.Sudo {
		args = append([]string{"-n", name}, args...)
		name = "sudo"
	}

	return &Operation{
		Execution: lexec.New(lexec.Loggerf(logger.Tracef), name, args...),
	}
}
