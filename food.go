package faces

import (
	"github.com/reconquest/faces/api/bash"
	"github.com/reconquest/faces/api/hastur"
)

func (context context) NewBash() (*bash.Bash, error) {
	face := new(bash.Bash)

	err := fabricate(face, context.Execution, "bash")
	if err != nil {
		return nil, err
	}

	return face, nil
}

func NewBash() (*bash.Bash, error) {
	return new(context).NewBash()
}

func (context context) NewHastur() (*hastur.Hastur, error) {
	face := new(hastur.Hastur)

	err := fabricate(face, context.Execution, "hastur")
	if err != nil {
		return nil, err
	}

	return face, nil
}

func NewHastur() (*hastur.Hastur, error) {
	return new(context).NewHastur()
}
