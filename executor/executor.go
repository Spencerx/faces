package executor

import (
	"github.com/reconquest/faces/logger"
	"github.com/reconquest/go-loggedexec"
)

type Executor struct {
	Sudo bool
}

func (executor *Executor) Command(
	name string,
	args ...string,
) *loggedexec.Execution {
	return loggedexec.New(loggedexec.Loggerf(logger.Tracef), name, args...)
}
