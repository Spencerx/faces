package faces

import (
	"github.com/reconquest/faces/commands/bash"
	"github.com/reconquest/faces/commands/true"
)

func (context context) NewBash() (*bash.Bash, error) {
	face := new(bash.Bash)

	err := fabricate(face, context.executor, "bash")
	if err != nil {
		return nil, err
	}

	return face, nil
}

func NewBash() (*bash.Bash, error) {
	return new(context).NewBash()
}

func (context context) NewTrue() (*true.True, error) {
	face := new(true.True)

	err := fabricate(face, context.executor, "true")
	if err != nil {
		return nil, err
	}

	return face, nil
}

func NewTrue() (*true.True, error) {
	return new(context).NewTrue()
}
