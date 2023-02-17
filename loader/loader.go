package loader

import (
	"fmt"
	"os"
	"os/exec"
	"pikman/alpine"
	"pikman/arch"
	"pikman/fedora"
	"pikman/flatpak"
	"pikman/types"
	"pikman/ubuntu"
	"strings"
)

var commandsMap = map[types.OSType]map[string]string{
	types.Ubuntu:  ubuntu.Commands,
	types.Arch:    arch.Commands,
	types.Fedora:  fedora.Commands,
	types.Alpine:  alpine.Commands,
	types.Flatpak: flatpak.Commands,
}

var packageManagerMap = map[types.OSType]string{
	types.Ubuntu:  ubuntu.PackageManager,
	types.Arch:    arch.PackageManager,
	types.Fedora:  fedora.PackageManager,
	types.Alpine:  alpine.PackageManager,
	types.Flatpak: flatpak.PackageManager,
}

func ProcessCommand(command string, osType types.OSType, containerName string, packageName []string, upgradableFlag bool) error {
	var err error
	if osType != types.Ubuntu && osType != types.Flatpak && containerName != "" {
		packageName = append([]string{"--name " + containerName}, packageName...)
	}

	if osType == types.Ubuntu && upgradableFlag {
		packageName = append([]string{"--upgradable"}, packageName...)
	}

	commandToExecute, err := getCommand(command, osType, packageName)
	cmd := exec.Command("/bin/sh", "-c", commandToExecute)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func getCommand(command string, osType types.OSType, packageName []string) (string, error) {
	commandMap := commandsMap[osType]
	cmd, ok := commandMap[command]

	if ok {
		return fmt.Sprintf("%s %s %s", packageManagerMap[osType], cmd, strings.Join(packageName, " ")), nil
	}

	return "", fmt.Errorf("%s: is not a valid command for this distro", command)
}
