package hastur

import "strings"

type Container struct {
	args map[string]string
}

func (container *Container) SetRootDirectory(directory string) {
	container.args["-r"] = directory
}

func (container *Container) SetPackages(packages []string) {
	container.args["-p"] = strings.Join(packages, ",")
}

func (container *Container) SetCommand(command string) {
	container.args["--"] = command
}

func (container *Container) SetHostNetwork(iface string) {
	container.args["-t"] = iface
}

func (container *Container) SetSourceDirectory(directory string) {
	container.args["-x"] = directory
}

func (container *Container) SetName(name string) {
	container.args["-n"] = name
}

func (container *Container) SetBridge(bridge string) {
	container.args["-b"] = bridge
}

func (container *Container) SetAddress(address string) {
	container.args["-a"] = address
}

func (container *Container) SetForce(force bool) {
	container.toggle("-f", force)
}

func (container *Container) SetQuietMode(quiet bool) {
	container.toggle("-q", quiet)
}

func (container *Container) SetKeep(keep bool) {
	container.toggle("-k", keep)
}

func (container *Container) SetKeepOnFail(keep bool) {
	container.toggle("-e", keep)
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
