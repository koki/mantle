package hostpath

type HostPathType string

const (
	HostPathUnset             HostPathType = ""
	HostPathDirectoryOrCreate HostPathType = "dir-or-create"
	HostPathDirectory         HostPathType = "dir"
	HostPathFileOrCreate      HostPathType = "file-or-create"
	HostPathFile              HostPathType = "file"
	HostPathSocket            HostPathType = "socket"
	HostPathCharDev           HostPathType = "char-dev"
	HostPathBlockDev          HostPathType = "block-dev"
)

type HostPathVolume struct {
	Path string       `json:"-"`
	Type HostPathType `json:"-"`
}
