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

func ProcessCommand(command string, osType types.OSType, containerName string, packageName []string) error {
	var err error
	if osType != types.Ubuntu && osType != types.Flatpak && containerName != "" {
		containerName = "--name " + containerName
	} else {
		containerName = ""
	}

	commandToExecute, err := getCommand(command, osType, containerName, packageName)
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

func getCommand(command string, osType types.OSType, containerName string, packageName []string) (string, error) {
	switch osType {
	case types.Arch:
		cmd, ok := arch.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", arch.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Arch", command)
		}
	case types.Fedora:
		cmd, ok := fedora.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", fedora.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Fedora", command)
		}
	case types.Flatpak:
		cmd, ok := flatpak.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s", flatpak.PackageManager, cmd, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Flatpak", command)
		}
	case types.Alpine:
		cmd, ok := alpine.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", alpine.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Alpine", command)
		}
	case types.Ubuntu:
		cmd, ok := ubuntu.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s", ubuntu.PackageManager, cmd, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Ubuntu", command)
		}
	}

	return "", fmt.Errorf("%s: was passed without a valid backend", command)
}
