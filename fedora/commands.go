package fedora

var PackageManager = "apx"

var Commands = map[string]string{
	"autoremove": "pikman-fedora-workstation autoremove",
	"clean":      "pikman-fedora-workstation clean",
	"export":     "pikman-fedora-workstation export",
	"init":       "subsystems new -n pikman-fedora-workstation -s fedora-workstation",
	"install":    "pikman-fedora-workstation install",
	"list":       "pikman-fedora-workstation list",
	"purge":      "pikman-fedora-workstation purge",
	"run":        "pikman-fedora-workstation run",
	"remove":     "pikman-fedora-workstation remove",
	"enter":      "pikman-fedora-workstation enter",
	"search":     "pikman-fedora-workstation search",
	"show":       "pikman-fedora-workstation show",
	"update":     "pikman-fedora-workstation update",
	"upgrade":    "pikman-fedora-workstation upgrade",
	"unexport":   "pikman-fedora-workstation unexport",
}
