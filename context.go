package faces

import "github.com/reconquest/faces/executor"

type context struct {
	executor executor.Executor
}

func (c *context) Sudo() *context {
	executor := c.executor

	executor.Sudo = true

	return &context{
		executor: executor,
	}
}
