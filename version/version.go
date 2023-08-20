package version

import (
	"fmt"
	"log"
	"runtime/debug"
)

var tag = "unknown"
var rev = "unknown"
var dirty = false

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("unable to read build info from binary")
		return
	}
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

func String() string {
	ver := fmt.Sprintf("%s-%s", tag, rev)
	if dirty {
		ver = ver + "-dirty"
	}
	return ver
}
