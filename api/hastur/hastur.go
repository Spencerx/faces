package hastur

import (
	"os/exec"
	"strings"

	"github.com/reconquest/faces/execution"
	"github.com/reconquest/faces/face"
	"github.com/reconquest/lexec-go"
)

var _ face.Interface = (*Hastur)(nil)

type Hastur struct {
	face.Abstract
	execution.Execution
}

func (hastur *Hastur) Init(execution execution.Execution) error {
	_, err := exec.LookPath("hastur")
	if err != nil {
		return err
	}

	hastur.Execution = execution

	return nil
}

func (hastur *Hastur) GetVersion() (string, error) {
	stdout, _, err := hastur.Exec("hastur", "--version").NoLog().Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

func (hastur *Hastur) NewContainer() *Container {
	return &Container{
		args: map[string]string{},
	}
}

func (hastur *Hastur) Start(
	container *Container,
) *lexec.Execution {
	args := []string{}
	for key, value := range container.args {
		if key == "--" {
			continue
		}

		args = append(args, key)
		if value != "" {
			args = append(args, value)
		}
	}

	args = append(args, "-S")

	if _, ok := container.args["--"]; ok {
		args = append(args, "--", container.args["--"])
	}

	return hastur.Exec("hastur", args...)
}
