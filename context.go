package faces

import "github.com/reconquest/faces/executor"

type context struct {
	executor executor.Executor
}

func Sudo() *context {
	context := new(context)

	context.executor.Sudo = true

	return context
}
