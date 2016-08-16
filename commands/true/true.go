package true

import (
	"errors"
	"os/exec"
	"regexp"

	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/executor"
	"github.com/reconquest/faces/face"
)

type True struct {
	logger lorg.Logger

	executor executor.Executor
}

var (
	_ face.Face = (*True)(nil)
)

func (true *True) Init(executor executor.Executor) error {
	_, err := exec.LookPath("true")
	if err != nil {
		return err
	}

	true.executor = executor

	return nil
}

func (true *True) GetVersion() (string, error) {
	stdout, _, err := true.executor.Command("true", "--version").Output()
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
