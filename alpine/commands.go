package alpine

var PackageManager = "apx"

var ContainerSubsystem = "pikman-alpine"

var ApxSubsystem = "alpine"

var Commands = map[string]string{
	"autoremove": "pikman-alpine autoremove",
	"clean":      "pikman-alpine clean",
	"export":     "pikman-alpine export",
	"init":       "subsystems new -n pikman-alpine -s alpine",
	"install":    "pikman-alpine install",
	"list":       "pikman-alpine list",
	"purge":      "pikman-alpine purge",
	"run":        "pikman-alpine run",
	"remove":     "pikman-alpine remove",
	"enter":      "pikman-alpine enter",
	"search":     "pikman-alpine search",
	"show":       "pikman-alpine show",
	"update":     "pikman-alpine update",
	"upgrade":    "pikman-alpine upgrade",
	"unexport":   "pikman-alpine unexport",
}
