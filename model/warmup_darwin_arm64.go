//go:build darwin && arm64

package model

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/sys/unix"
)

var pageSize = unix.Getpagesize()

const slcTargetBytes = 32 * 1024 * 1024

func warmUpWeights(f *os.File, data []byte) {
	defer f.Close()

	if len(data) == 0 {
		return
	}

	if os.Getenv("OLLAMA_DISABLE_WARMUP") == "1" {
		return
	}

	budget := slcTargetBytes
	if s := os.Getenv("OLLAMA_WARMUP_BYTES"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			budget = n
		}
	}
	if budget > len(data) {
		budget = len(data)
	}
	if budget == 0 {
		return
	}

	start := time.Now()
	deadline := start.Add(100 * time.Millisecond)

	_ = unix.Madvise(data[:budget], unix.MADV_WILLNEED)
	locked := false
	if err := unix.Mlock(data[:budget]); err == nil {
		locked = true
	}

	for off := 0; off < budget; off += pageSize {
		_ = data[off]
		if time.Now().After(deadline) {
			break
		}
	}

	if locked {
		_ = unix.Munlock(data[:budget])
	}

	if os.Getenv("OLLAMA_WARMUP_VERBOSE") == "1" {
		log.Printf("[warmup] touched=%dMB in %v", budget/1024/1024, time.Since(start))
	}

	runtime.KeepAlive(data)
	_ = unix.Munmap(data)
}
