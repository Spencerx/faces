package execution

import (
	"github.com/reconquest/faces/logger"
	"github.com/reconquest/loggedexec-go"
)

type Execution struct {
	Sudo bool
}

func (execution *Execution) Exec(
	name string, args ...string,
) *loggedexec.Execution {
	if execution.Sudo {
		args = append([]string{"-n", name}, args...)
		name = "sudo"
	}

	return loggedexec.New(loggedexec.Loggerf(logger.Tracef), name, args...)
}
