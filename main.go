package main

import (
	"log"
	"os"
	"pikman/loader"
	"pikman/types"

	"github.com/urfave/cli/v2"
)

func main() {

	if os.Getuid() == 0 {
		log.Fatalf("Error: Do not run pikman as root")
	}

	osType := types.Ubuntu
	containerName := ""
	upgradableFlag := false
	installedFlag := false
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version number",
	}

	app := &cli.App{
		Name:                 "pikman",
		Usage:                "One package manager to rule them all",
		Version:              "v1.23.2.17.0",
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
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"cl"},
				Usage:   "Clean the package manager cache",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "enter",
				Usage: "Enter the container instance for select package manager",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "export",
				Usage: "Export/Recreate a program's desktop entry from the container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "init",
				Usage: "Initialize a managed container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install the specified package(s)",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List installed packages",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "upgradable",
						Usage:       "Used by list to check upgradable packages",
						Destination: &upgradableFlag,
					},
					&cli.BoolFlag{
						Name:        "installed",
						Usage:       "Used by list to check installed packages",
						Destination: &installedFlag,
					},
				},
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), upgradableFlag, installedFlag)
				},
			},
			{
				Name:  "log",
				Usage: "Show package manager logs",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "purge",
				Usage: "Fully purge a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "run",
				Usage: "Run a command inside a managed container",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove an installed package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Search for a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "show",
				Usage: "Show details for a package",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "unexport",
				Usage: "Unexport/Remove a program's desktop entry",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "update",
				Usage: "Update the list of available packages",
				Action: func(cCtx *cli.Context) error {
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
			{
				Name:  "upgrade",
				Usage: "Upgrade the system by installing/upgrading available packages",
				Action: func(cCtx *cli.Context) error {
					cCtx.Args().Tail()
					return loader.ProcessCommand(cCtx.Command.FullName(), osType, containerName, cCtx.Args().Slice(), false, false)
				},
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
