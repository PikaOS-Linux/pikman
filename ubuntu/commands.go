package ubuntu

var PackageManager = "sudo -S nala"

var Commands = map[string]string{
	"autoremove": "autoremove",
	"clean":      "clean",
	"install":    "install",
	"list":       "list",
	"purge":      "purge",
	"remove":     "remove",
	"search":     "search",
	"show":       "show",
	"update":     "update",
	"upgrade":    "upgrade",
}
