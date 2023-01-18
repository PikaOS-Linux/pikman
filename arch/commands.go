package arch

var PackageManager = "apx"

var Commands = map[string]string{
	"autoremove": "--aur autoremove",
	"clean":      "--aur clean",
	"export":     "--aur export",
	"init":       "--aur init",
	"install":    "--aur install",
	"list":       "--aur list",
	"log":        "--aur log",
	"purge":      "--aur purge",
	"run":        "--aur run",
	"remove":     "--aur remove",
	"enter":      "--aur enter",
	"search":     "--aur search",
	"show":       "--aur show",
	"update":     "--aur update",
	"upgrade":    "--aur upgrade",
}
