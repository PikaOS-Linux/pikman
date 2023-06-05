package main

import (
	"log"
	"os"

	"pikman/command"
	"pikman/types"

	"github.com/urfave/cli/v2"
)

func main() {
	if os.Getuid() == 0 {
		log.Fatalf("Error: Do not run pikman as root")
	}

	cmd := &command.Command{
		OsType:        types.Ubuntu,
		ContainerName: "",
		IsUpgradable:  false,
		IsJSON:        false,
		IsInstalled:   false,
		PackageName:   make([]string, 0),
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version number",
	}

	app := &cli.App{
		Name:                 "pikman",
		Usage:                "One package manager to rule them all",
		Version:              "v1.23.6.5.2",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "arch",
				Aliases: []string{"aur"},
				Usage:   "Install Arch packages (including from the AUR)",
				Action: func(cCtx *cli.Context, b bool) error {
					if b {
						cmd.OsType = types.Arch
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
						cmd.OsType = types.Fedora
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
						cmd.OsType = types.Alpine
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
						cmd.OsType = types.Flatpak
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name of the managed container",
				Destination: &cmd.ContainerName,
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "autoremove",
				Usage:  "Remove all unused packages",
				Action: cmd.Process,
			},
			{
				Name:    "clean",
				Aliases: []string{"cl"},
				Usage:   "Clean the package manager cache",
				Action:  cmd.Process,
			},
			{
				Name:   "enter",
				Usage:  "Enter the container instance for select package manager",
				Action: cmd.Process,
			},
			{
				Name:   "export",
				Usage:  "Export/Recreate a program's desktop entry from the container",
				Action: cmd.Process,
			},
			{
				Name:   "init",
				Usage:  "Initialize a managed container",
				Action: cmd.Process,
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install the specified package(s)",
				Action:  cmd.Process,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List installed packages",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "upgradable",
						Usage:       "Used by list to check upgradable packages",
						Destination: &cmd.IsUpgradable,
					},
					&cli.BoolFlag{
						Name:        "installed",
						Usage:       "Used by list to check installed packages",
						Destination: &cmd.IsInstalled,
					},
				},
				Action: cmd.Process,
			},
			{
				Name:   "log",
				Usage:  "Show package manager logs",
				Action: cmd.Process,
			},
			{
				Name:   "purge",
				Usage:  "Fully purge a package",
				Action: cmd.Process,
			},
			{
				Name:   "run",
				Usage:  "Run a command inside a managed container",
				Action: cmd.Process,
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove an installed package",
				Action:  cmd.Process,
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Search for a package",
				Action:  cmd.Process,
			},
			{
				Name:   "show",
				Usage:  "Show details for a package",
				Action: cmd.Process,
			},
			{
				Name:   "unexport",
				Usage:  "Unexport/Remove a program's desktop entry",
				Action: cmd.Process,
			},
			{
				Name:   "update",
				Usage:  "Update the list of available packages",
				Action: cmd.Process,
			},
			{
				Name:   "upgrade",
				Usage:  "Upgrade the system by installing/upgrading available packages",
				Action: cmd.Process,
			},
			{
				Name:  "upgrades",
				Usage: "List the available upgrades",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "json",
						Usage:       "Output in JSON",
						Destination: &cmd.IsJSON,
					},
				},
				Action: cmd.Process,
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
