//go:build linux && amd64

package model

import (
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

func warmUpWeights(_ *os.File, data []byte) {
	if len(data) == 0 {
		return
	}
	_ = unix.Madvise(data, unix.MADV_POPULATE_READ)
	runtime.KeepAlive(data)
	_ = unix.Munmap(data)
}
