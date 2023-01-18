package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"pikman/alpine"
	"pikman/arch"
	"pikman/fedora"
	"pikman/flatpak"
	"pikman/ubuntu"
	"strings"
)

type OSType = int

const (
	Ubuntu OSType = iota
	Arch
	Fedora
	Alpine
	Flatpak
)

func main() {

	osType := Ubuntu
	containerName := ""

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version number",
	}

	app := &cli.App{
		Name:                 "pikman",
		Usage:                "One package manager to rule them all",
		Version:              "v0.11",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "arch",
				Aliases: []string{"aur"},
				Usage:   "Install Arch packages (including from the AUR)",
				Action: func(cCtx *cli.Context, b bool) error {
					if b {
						osType = Arch
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "fedora",
				Aliases: []string{"dnf"},
				Usage:   "Install Fedora packages",
				Action: func(cCtx *cli.Context, b bool) error {
					if b {
						osType = Fedora
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "alpine",
				Aliases: []string{"apk"},
				Usage:   "Install Alpine packages",
				Action: func(cCtx *cli.Context, b bool) error {
					if b {
						osType = Alpine
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "flatpak",
				Aliases: []string{"fl"},
				Usage:   "Install Flatpak packages",
				Action: func(cCtx *cli.Context, b bool) error {
					if b {
						osType = Flatpak
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name of the managed container",
				Destination: &containerName,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "autoremove",
				Usage: "Remove all unused packages",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"cl"},
				Usage:   "Clean the package manager cache",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "enter",
				Usage: "Enter the container instance for select package manager",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "export",
				Usage: "Export/Recreate a program's desktop entry from the container",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "init",
				Usage: "Initialize a managed container",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install the specified package(s)",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List installed packages",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "log",
				Usage: "Show package manager logs",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "purge",
				Usage: "Fully purge a package",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "run",
				Usage: "Run a command inside a managed container",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove an installed package",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Search for a package",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "show",
				Usage: "Show details for a package",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "unexport",
				Usage: "Unexport/Remove a program's desktop entry",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "update",
				Usage: "Update the list of available packages",
				Action: func(cCtx *cli.Context) error {
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "upgrade",
				Usage: "Upgrade the system by installing/upgrading available packages",
				Action: func(cCtx *cli.Context) error {
					cCtx.Args().Tail()
					return processCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func processCommand(command string, osType OSType, containerName string, packageName []string) error {
	var err error
	if osType == Arch && containerName != "" {
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

func getCommand(command string, osType OSType, containerName string, packageName []string) (string, error) {
	switch osType {
	case Arch:
		cmd, ok := arch.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", arch.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Arch", command)
		}
	case Fedora:
		cmd, ok := fedora.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", fedora.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Fedora", command)
		}
	case Flatpak:
		cmd, ok := flatpak.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s", flatpak.PackageManager, cmd, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Flatpak", command)
		}
	case Alpine:
		cmd, ok := alpine.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s %s", alpine.PackageManager, cmd, containerName, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Alpine", command)
		}
	case Ubuntu:
		cmd, ok := ubuntu.Commands[command]
		if ok {
			return fmt.Sprintf("%s %s %s", ubuntu.PackageManager, cmd, strings.Join(packageName, " ")), nil
		} else {
			return "", fmt.Errorf("%s: is not a valid command for Ubuntu", command)
		}
	}

	return "", fmt.Errorf("%s: was passed without a valid backend", command)
}
