package true

import (
	"errors"
	"os/exec"
	"regexp"

	"github.com/kovetskiy/executil"
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces"
)

type True struct {
	logger lorg.Logger
}

var (
	_ faces.Face = (*True)(nil)
)

func (*True) Init() error {
	_, err := exec.LookPath("true")
	if err != nil {
		return err
	}

	return nil
}

func (*True) GetVersion() (string, error) {
	stdout, _, err := executil.Run(exec.Command("true", "--version"))
	if err != nil {
		return "", err
	}

	// true (GNU coreutils) 8.25
	// Copyright (C) 2016 Free Software Foundation, Inc.
	// License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
	// This is free software: you are free to change and redistribute it.
	// There is NO WARRANTY, to the extent permitted by law.
	version := regexp.MustCompile(`true \(GNU coreutils\) (.*)`).Find(stdout)
	if len(version) == 0 {
		return "", errors.New("ambiguous output without version definition")
	}

	return string(version), nil
}

func (true *True) SetLogger(logger lorg.Logger) {
	true.logger = logger
}
