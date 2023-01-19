package types

type OSType = int

const (
	Ubuntu OSType = iota
	Arch
	Fedora
	Alpine
	Flatpak
)
