package faces

import "github.com/seletskiy/hierr"

func fabricate(face Face, name string) error {
	face.SetLogger(logger)

	err := face.Init()
	if err != nil {
		return hierr.Errorf(err, "can't initialize %s", name)
	}

	version, err := face.GetVersion()
	if err != nil {
		return hierr.Errorf(err, "can't obtain version of %s", name)
	}

	debugf("%s version: %s", name, version)

	return nil
}
