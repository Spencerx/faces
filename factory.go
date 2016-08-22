package faces

import (
	"github.com/reconquest/faces/execution"
	"github.com/reconquest/faces/face"
	"github.com/reconquest/faces/logger"
	"github.com/reconquest/ser-go"
)

func fabricate(
	face face.Interface,
	exec execution.Execution,
	name string,
) error {
	face.SetLogger(logger.Logger)

	err := face.Init(exec)
	if err != nil {
		return ser.Errorf(err, "can't initialize %s", name)
	}

	version, err := face.GetVersion()
	if err != nil {
		return ser.Errorf(err, "can't obtain version of %s", name)
	}

	logger.Debugf("%s version: %s", name, version)

	return nil
}
