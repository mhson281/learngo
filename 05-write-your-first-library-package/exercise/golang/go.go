package golang

import (
	"runtime"
)

func Version() {
	return runtime.Version()
}
