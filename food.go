package faces

import (
	"github.com/reconquest/faces/true"
)

func NewTrue() (*true.True, error) {
	face := new(true.True)

	err := fabricate(face, "true")
	if err != nil {
		return nil, err
	}

	return face, nil
}
