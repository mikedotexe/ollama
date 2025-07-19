//go:build windows

package model

import "os"

func mmap(_ *os.File) ([]byte, error) { return nil, nil }
