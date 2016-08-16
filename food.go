package faces

import (
	"github.com/reconquest/faces/commands/fdisk"
	"github.com/reconquest/faces/commands/true"
)

func (context context) NewFdisk() (*fdisk.Fdisk, error) {
	face := new(fdisk.Fdisk)

	err := fabricate(face, context.executor, "fdisk")
	if err != nil {
		return nil, err
	}

	return face, nil
}

func NewFdisk() (*fdisk.Fdisk, error) {
	return new(context).NewFdisk()
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
