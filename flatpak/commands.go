package flatpak

var PackageManager = "flatpak"

var Commands = map[string]string{
	"install": "install",
	"list":    "list",
	"log":     "history",
	"run":     "run",
	"remove":  "uninstall",
	"enter":   "enter",
	"search":  "search",
	"show":    "info",
	"update":  "update",
}
