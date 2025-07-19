//go:build !darwin || !arm64

package model

import "os"

func warmUpWeights(_ *os.File, _ []byte) {}
