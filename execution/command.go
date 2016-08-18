package execution

import (
	"bytes"
	"errors"
	"io"

	"github.com/kovetskiy/executil"
	"github.com/reconquest/faces/logger"
	"github.com/reconquest/go-nopio"
	"github.com/reconquest/go-prefixwriter"
	"github.com/reconquest/lexec-go"
	"github.com/seletskiy/hierr"
)

type Execution struct {
	Sudo bool
}

func (execution *Execution) Exec(
	name string, args ...string,
) *lexec.Execution {
	if execution.Sudo {
		args = append([]string{"-n", name}, args...)
		name = "sudo"
	}

	return lexec.New(lexec.Loggerf(logger.Tracef), name, args...)
}

func (execution *Execution) Errorf(
	stdout []byte,
	stderr []byte,
	err error,
	format string,
	args ...interface{},
) error {
	if !executil.IsExitError(err) {
		return err
	}

	buffer := &bytes.Buffer{}

	output := &struct {
		io.Writer
		io.Closer
	}{
		Writer: buffer,
		Closer: nopio.NopCloser{},
	}

	_, _ = prefixwriter.New(output, `{stdout} `).Write(stdout)
	_, _ = prefixwriter.New(output, `{stderr} `).Write(stderr)

	return hierr.Errorf(
		errors.New(buffer.String()),
		format+` (exit code %d)`,
		append(args, executil.GetExitStatus(err))...,
	)
}
