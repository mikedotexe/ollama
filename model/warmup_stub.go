//go:build !(darwin && arm64) && !(linux && amd64)

package model

import "os"

func warmUpWeights(_ *os.File, _ []byte) {}
