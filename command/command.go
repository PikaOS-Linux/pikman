package command

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"pikman/alpine"
	"pikman/arch"
	"pikman/fedora"
	"pikman/flatpak"
	"pikman/types"
	"pikman/ubuntu"
	"pikman/updates"

	"github.com/urfave/cli/v2"
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

var containerSubsystemMap = map[types.OSType]string{
	types.Arch:   arch.ContainerSubsystem,
	types.Fedora: fedora.ContainerSubsystem,
	types.Alpine: alpine.ContainerSubsystem,
}

func (c *Command) ContainerSubsystem() string {
	return containerSubsystemMap[c.OsType]
}

var apxSubsystemMap = map[types.OSType]string{
	types.Arch:   arch.ApxSubsystem,
	types.Fedora: fedora.ApxSubsystem,
	types.Alpine: alpine.ApxSubsystem,
}

func (c *Command) ApxSubsystem() string {
	return apxSubsystemMap[c.OsType]
}

type Command struct {
	Command       string
	OsType        types.OSType
	ContainerName string
	PackageName   []string
	IsUpgradable  bool
	IsInstalled   bool
	IsJSON        bool
}

func (c *Command) Process(cCtx *cli.Context) error {
	c.Command = cCtx.Command.FullName()
	c.PackageName = cCtx.Args().Slice()
	return c.processCommand()
}

func (c *Command) processCommand() error {
	if (c.OsType == types.Ubuntu || c.OsType == types.Flatpak) && c.Command == "upgrades" {
		return c.runUpgrades()
	}

	var err error

	if c.OsType != types.Ubuntu && c.OsType != types.Flatpak && c.Command != "init" {
		err = c.initContainer()
		if err != nil {
			return err
		}
	}

	if c.OsType != types.Ubuntu && c.OsType != types.Flatpak && c.ContainerName != "" {
		c.PackageName = append([]string{"--name " + c.ContainerName}, c.PackageName...)
	}

	if c.OsType == types.Ubuntu && c.IsUpgradable {
		c.PackageName = append([]string{"--upgradable"}, c.PackageName...)
	}

	if c.OsType == types.Ubuntu && c.IsInstalled {
		c.PackageName = append([]string{"--installed"}, c.PackageName...)
	}

	commandToExecute, err := c.getCommand()
	if err != nil {
		return err
	}

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

func (c *Command) getCommand() (string, error) {
	commandMap := commandsMap[c.OsType]
	cmd, ok := commandMap[c.Command]

	if ok {
		return fmt.Sprintf("%s %s %s", packageManagerMap[c.OsType], cmd, strings.Join(c.PackageName, " ")), nil
	}

	return "", fmt.Errorf("%s: is not a valid command for this distro", c.Command)
}

func (c *Command) runUpgrades() error {
	if c.OsType == types.Ubuntu {
		return updates.GetUbuntuUpdates(c.IsJSON)
	}

	return updates.GetFlatpakUpdates(c.IsJSON)
}

func (c *Command) initContainer() error {
	contName := c.ContainerName
	if contName == "" {
		contName = c.ContainerSubsystem()
	}

	cmd := exec.Command("/bin/bash", "-c", "apx subsystems list -j")
	outb, err := cmd.Output()
	if err != nil {
		return err
	}

	var out []apxSubs
	err = json.Unmarshal(outb, &out)
	if err != nil {
		return err
	}

	isMatch := false
	for _, sub := range out {
		if sub.Name == contName {
			isMatch = true
			break
		}
	}

	if !isMatch {
		fmt.Println("Warning: Subsystem hasn't been pre-initialized, initializing...")
		cmd_exec := exec.Command("/bin/bash", "-c", "apx subsystems new -n "+contName+" -s "+c.ApxSubsystem())
		cmd_exec.Stdout = os.Stdout
		cmd_exec.Stdin = os.Stdin
		cmd_exec.Stderr = os.Stderr
		if err := cmd_exec.Run(); err != nil {
			fmt.Println("Apx Error: ", err)
			return err
		}
	}
	return nil
}

type apxSubs struct {
	Name string `json:"name"`
}
