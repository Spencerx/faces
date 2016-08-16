package faces

import (
	"github.com/reconquest/faces/executor"
	"github.com/reconquest/faces/face"
	"github.com/reconquest/faces/logger"
	"github.com/seletskiy/hierr"
)

func fabricate(
	face face.Face,
	executor executor.Executor,
	name string,
) error {
	face.SetLogger(logger.Logger)

	err := face.Init(executor)
	if err != nil {
		return hierr.Errorf(err, "can't initialize %s", name)
	}

	version, err := face.GetVersion()
	if err != nil {
		return hierr.Errorf(err, "can't obtain version of %s", name)
	}

	logger.Debugf("%s version: %s", name, version)

	return nil
}
