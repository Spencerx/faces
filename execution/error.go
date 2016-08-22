package execution

import (
	"bytes"
	"errors"
	"io"

	"github.com/kovetskiy/executil"
	"github.com/reconquest/lexec-go"
	"github.com/reconquest/nopio-go"
	"github.com/reconquest/prefixwriter-go"
	"github.com/reconquest/ser-go"
)

func expandError(
	err error, streamsData []lexec.StreamData,
) error {
	if !executil.IsExitError(err) {
		return err
	}

	execError, ok := err.(*executil.Error)
	if !ok {
		return err
	}

	var (
		buffer = &bytes.Buffer{}

		bufferWriter = &struct {
			io.Writer
			io.Closer
		}{
			Writer: buffer,
			Closer: nopio.NopCloser{},
		}

		stdoutWriter = prefixwriter.New(bufferWriter, `{stdout} `)
		stderrWriter = prefixwriter.New(bufferWriter, `{stderr} `)
	)

	for _, streamData := range streamsData {
		switch streamData.Stream {
		case lexec.Stdout:
			_, _ = stdoutWriter.Write(streamData.Data)
		case lexec.Stderr:
			_, _ = stderrWriter.Write(streamData.Data)
		}
	}

	return ser.Errorf(
		errors.New(buffer.String()),
		"%q (exit code %d)",
		execError.Cmd.Args, executil.GetExitStatus(execError),
	)
}
