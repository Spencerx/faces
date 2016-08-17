package bash

import (
	"errors"
	"os/exec"
	"regexp"

	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/executor"
	"github.com/reconquest/faces/face"
)

type Bash struct {
	logger   lorg.Logger
	executor executor.Executor
}

var (
	_ face.Face = (*Bash)(nil)
)

func (bash *Bash) Init(executor executor.Executor) error {
	_, err := exec.LookPath("bash")
	if err != nil {
		return err
	}

	bash.executor = executor

	return nil
}

func (bash *Bash) GetVersion() (string, error) {
	stdout, _, err := bash.executor.Command(
		"bash", "--version",
	).NoLog().Output()
	if err != nil {
		return "", err
	}

	// bash (GNU coreutils) 8.25
	// Copyright (C) 2016 Free Software Foundation, Inc.
	// License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
	// This is free software: you are free to change and redistribute it.
	// There is NO WARRANTY, to the extent permitted by law.
	version := regexp.MustCompile(`version (.*)`).Find(stdout)
	if len(version) == 0 {
		return "", errors.New("ambiguous output without version definition")
	}

	return string(version), nil
}

func (bash *Bash) SetLogger(logger lorg.Logger) {
	bash.logger = logger
}

func (bash *Bash) Eval(
	code string,
) (stdout []byte, stderr []byte, err error) {
	return bash.executor.Command("bash", "-c", code).Output()
}
