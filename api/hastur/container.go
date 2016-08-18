package hastur

import "strings"

type Container struct {
	args map[string]string
}

func (container *Container) SetRootDirectory(directory string) *Container {
	container.args["-r"] = directory
	return container
}

func (container *Container) SetPackages(packages []string) *Container {
	container.args["-p"] = strings.Join(packages, ",")
	return container
}

func (container *Container) SetCommand(command string) *Container {
	container.args["--"] = command
	return container
}

func (container *Container) SetHostNetwork(iface string) *Container {
	container.args["-t"] = iface
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
	container.toggle("-f", force)
	return container
}

func (container *Container) SetQuietMode(quiet bool) *Container {
	container.toggle("-q", quiet)
	return container
}

func (container *Container) SetKeep(keep bool) *Container {
	container.toggle("-k", keep)
	return container
}

func (container *Container) SetKeepOnFail(keep bool) *Container {
	container.toggle("-e", keep)
	return container
}
func (container *Container) toggle(key string, value bool) {
	if value {
		container.args[key] = ""
	} else {
		if _, ok := container.args[key]; ok {
			delete(container.args, key)
		}
	}
}
