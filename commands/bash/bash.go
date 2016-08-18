package bash

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/kovetskiy/lorg"
	"github.com/reconquest/faces/execution"
	"github.com/reconquest/faces/face"
)

var _ face.Face = (*Bash)(nil)

type Bash struct {
	lorg.Logger
	execution.Execution
}

func (bash *Bash) Init(execution execution.Execution) error {
	_, err := exec.LookPath("bash")
	if err != nil {
		return err
	}

	bash.Execution = execution

	return nil
}

func (bash *Bash) GetVersion() (string, error) {
	stdout, _, err := bash.Exec(
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

	return strings.Fields(string(version))[1], nil
}

func (bash *Bash) SetLogger(logger lorg.Logger) {
	bash.Logger = logger
}

func (bash *Bash) Eval(
	code string,
) (stdout []byte, stderr []byte, err error) {
	return bash.Exec("bash", "-c", code).Output()
}
