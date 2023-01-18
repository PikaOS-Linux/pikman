package alpine

var PackageManager = "apx "

var Commands = map[string]string{
	"autoremove": "--apk autoremove",
	"clean":      "--apk clean",
	"export":     "--apk export",
	"init":       "--apk init",
	"install":    "--apk install",
	"list":       "--apk list",
	"log":        "--apk log",
	"purge":      "--apk purge",
	"run":        "--apk run",
	"remove":     "--apk remove",
	"enter":      "--apk enter",
	"search":     "--apk search",
	"show":       "--apk show",
	"update":     "--apk update",
	"upgrade":    "--apk upgrade",
}
