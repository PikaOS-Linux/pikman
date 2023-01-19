package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"pikman/loader"
	"pikman/types"
)

func main() {

	osType := types.Ubuntu
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
						osType = types.Arch
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
						osType = types.Fedora
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
						osType = types.Alpine
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
						osType = types.Flatpak
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
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"cl"},
				Usage:   "Clean the package manager cache",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "enter",
				Usage: "Enter the container instance for select package manager",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "export",
				Usage: "Export/Recreate a program's desktop entry from the container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "init",
				Usage: "Initialize a managed container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install the specified package(s)",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List installed packages",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "log",
				Usage: "Show package manager logs",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "purge",
				Usage: "Fully purge a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "run",
				Usage: "Run a command inside a managed container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove an installed package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Search for a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "show",
				Usage: "Show details for a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "unexport",
				Usage: "Unexport/Remove a program's desktop entry",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "update",
				Usage: "Update the list of available packages",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
			{
				Name:  "upgrade",
				Usage: "Upgrade the system by installing/upgrading available packages",
				Action: func(cCtx *cli.Context) error {
					cCtx.Args().Tail()
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice())
				},
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
