package hastur

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/reconquest/faces/execution"
	"github.com/reconquest/faces/face"
	"github.com/reconquest/ser-go"
)

var _ face.Interface = (*Hastur)(nil)

type Hastur struct {
	face.Abstract
	execution.Execution

	rootDirectory string
	hostNetwork   string
	beQuiet       bool
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

func (hastur *Hastur) SetRootDirectory(directory string) *Hastur {
	hastur.rootDirectory = directory
	return hastur
}

func (hastur *Hastur) SetHostNetwork(network string) *Hastur {
	hastur.hostNetwork = network
	return hastur
}

func (hastur *Hastur) SetQuietMode(quiet bool) *Hastur {
	hastur.beQuiet = quiet
	return hastur
}

func (hastur *Hastur) Start(
	container *Container,
) *execution.Operation {
	args := []string{}

	if len(hastur.rootDirectory) > 0 {
		args = append(args, "-r", hastur.rootDirectory)
	}

	if len(hastur.hostNetwork) > 0 {
		args = append(args, "-t", hastur.hostNetwork)
	}

	if hastur.beQuiet {
		args = append(args, "-q")
	}

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

func (hastur *Hastur) Query(name ...string) ([]Container, error) {
	args := []string{}

	if len(hastur.rootDirectory) > 0 {
		args = append(args, "-r", hastur.rootDirectory)
	}

	args = append(args, "-j", "-Q")
	args = append(args, name...)

	stdout, _, err := hastur.Exec("hastur", args...).Output()
	if err != nil {
		return nil, err
	}

	var containers []Container
	err = json.Unmarshal(stdout, &containers)
	if err != nil {
		return nil, ser.Errorf(
			err, "can't unmarshal hastur output",
		)
	}

	return containers, nil
}
