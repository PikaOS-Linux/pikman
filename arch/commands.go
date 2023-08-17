package arch

var PackageManager = "apx"

var ContainerSubsystem = "pikman-arch-linux"

var ApxSubsystem = "arch-linux"

var Commands = map[string]string{
	"autoremove": "pikman-arch-linux autoremove",
	"clean":      "pikman-arch-linux clean",
	"export":     "pikman-arch-linux export",
	"init":       "subsystems new -n pikman-arch-linux -s arch-linux",
	"install":    "pikman-arch-linux install",
	"list":       "pikman-arch-linux list",
	"purge":      "pikman-arch-linux purge",
	"run":        "pikman-arch-linux run",
	"remove":     "pikman-arch-linux remove",
	"enter":      "pikman-arch-linux enter",
	"search":     "pikman-arch-linux search",
	"show":       "pikman-arch-linux show",
	"update":     "pikman-arch-linux update",
	"upgrade":    "pikman-arch-linux upgrade",
	"unexport":   "pikman-arch-linux unexport",
}
