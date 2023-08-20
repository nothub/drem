package buildinfo

import (
	"fmt"
	"log"
	"path"
	"runtime/debug"
)

var name = "unknown"
var module = "unknown"

var tag = "unknown"
var rev = "unknown"
var dirty = false

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("unable to read build info from binary")
		return
	}

	name = path.Base(bi.Main.Path)
	module = bi.Main.Path

	for _, bs := range bi.Settings {
		switch bs.Key {
		case "vcs.revision":
			rev = bs.Value[:8]
		case "vcs.modified":
			if bs.Value == "true" {
				dirty = true
			}
		}
	}
}

func Name() string {
	return name
}

func Module() string {
	return module
}

func Version() string {
	ver := fmt.Sprintf("%s-%s", tag, rev)
	if dirty {
		ver = ver + "-dirty"
	}
	return ver
}
