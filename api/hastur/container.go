package hastur

import "strings"

type Container struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Root    string `json:"root"`
	Address string `json:"address"`
	args    map[string]string
}

func (container *Container) SetPackages(packages []string) *Container {
	container.args["-p"] = strings.Join(packages, ",")
	return container
}

func (container *Container) SetCommand(command string) *Container {
	container.args["--"] = command
	return container
}

func (container *Container) SetSourceDirectory(directory string) *Container {
	container.args["-x"] = directory
	return container
}

func (container *Container) SetName(name string) *Container {
	container.args["-n"] = name
	return container
}

func (container *Container) SetBridge(bridge string) *Container {
	container.args["-b"] = bridge
	return container
}

func (container *Container) SetAddress(address string) *Container {
	container.args["-a"] = address
	return container
}

func (container *Container) SetForce(force bool) *Container {
	toggle(container.args, "-f", force)
	return container
}

func (container *Container) SetQuietMode(quiet bool) *Container {
	toggle(container.args, "-q", quiet)
	return container
}

func (container *Container) SetKeep(keep bool) *Container {
	toggle(container.args, "-k", keep)
	return container
}

func (container *Container) SetKeepOnFail(keep bool) *Container {
	toggle(container.args, "-e", keep)
	return container
}

func toggle(args map[string]string, arg string, value bool) {
	if value {
		args[arg] = ""
	} else {
		if _, ok := args[arg]; ok {
			delete(args, arg)
		}
	}
}
