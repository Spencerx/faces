package execution

import (
	"os/exec"

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

	command := exec.Command(name, args...)

	return &Operation{
		Execution: lexec.NewExec(lexec.Loggerf(logger.Tracef), command),
	}
}
