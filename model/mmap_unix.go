//go:build !windows

package model

import (
	"golang.org/x/sys/unix"
	"os"
)

func mmap(f *os.File) ([]byte, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := int(stat.Size())
	if size == 0 {
		return nil, nil
	}
	return unix.Mmap(int(f.Fd()), 0, size, unix.PROT_READ, unix.MAP_SHARED)
}
