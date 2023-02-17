# pikman - One package manager to rule them all

## USAGE:
pikman [global options] command [command options] [arguments...]

## VERSION:
v1.23.2.17.0

## COMMANDS:

- **autoremove** - Remove all unused packages
- **clean, cl** - Clean the package manager cache
- **enter** -     Enter the container instance for select package manager
- **export** -    Export/Recreate a program's desktop entry from the container
- **init** -      Initialize a managed container
- **install, i** -  Install the specified package(s)
- **list, l** -   List installed packages
- **log** -       Show package manager logs
- **purge** -     Fully purge a package
- **run** -       Run a command inside a managed container
- **remove, r** - Remove an installed package
- **search, s** - Search for a package
- **show**      - Show details for a package
- **unexport**  - Unexport/Remove a program's desktop entry
- **update**    - Update the list of available packages
- **upgrade**   - Upgrade the system by installing/upgrading available packages
- **help, h**   - Shows a list of commands or help for one command

## GLOBAL OPTIONS:

- **--arch, --aur** -  Install Arch packages (including from the AUR) (default: false)
- **--fedora, --dnf** - Install Fedora packages (default: false)
- **--alpine, --apk** - Install Alpine packages (default: false)
- **--flatpak, --fl** - Install Flatpak packages (default: false)
- **--name value** - Name of the managed container
- **--help, -h** - show help (default: false)
- **--version, -v** - Version number (default: false)