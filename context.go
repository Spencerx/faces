package faces

import "github.com/reconquest/faces/execution"

type context struct {
	execution.Execution
}

func Sudo() *context {
	context := new(context)

	context.Sudo = true

	return context
}
