package fedora

var PackageManager = "apx"

var Commands = map[string]string{
	"autoremove": "--dnf autoremove",
	"clean":      "--dnf clean",
	"export":     "--dnf export",
	"init":       "--dnf init",
	"install":    "--dnf install",
	"list":       "--dnf list",
	"log":        "--dnf log",
	"purge":      "--dnf purge",
	"run":        "--dnf run",
	"remove":     "--dnf remove",
	"enter":      "--dnf enter",
	"search":     "--dnf search",
	"show":       "--dnf show",
	"update":     "--dnf update",
	"upgrade":    "--dnf upgrade",
}
