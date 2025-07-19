//go:build darwin && arm64

package model

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/sys/unix"

	"github.com/ollama/ollama/internal/hardware"
)

var pageSize = unix.Getpagesize()

var slcTargetBytes = func() int {
	n := hardware.DetectSLCMiB() * 1024 * 1024
	if n == 0 {
		n = 32 * 1024 * 1024
	}
	return n
}()

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
	dur := 100 * time.Millisecond
	if s := os.Getenv("OLLAMA_WARMUP_BUDGET_MS"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n >= 0 {
			dur = time.Duration(n) * time.Millisecond
		}
	}
	if dur == 0 {
		return
	}
	deadline := start.Add(dur)

	_ = unix.Madvise(data[:budget], unix.MADV_WILLNEED)
	locked := unix.Mlock(data[:budget]) == nil
	if locked {
		defer unix.Munlock(data[:budget])
	}

	stride := 3 * pageSize
	for off := 0; off < budget; off += stride {
		_ = data[off]
		if time.Now().After(deadline) {
			break
		}
	}

	if locked {
		// unlock handled by defer
	}

	if os.Getenv("OLLAMA_WARMUP_VERBOSE") == "1" {
		log.Printf("[warmup] touched=%dMB in %v", budget/1024/1024, time.Since(start))
	}

	runtime.KeepAlive(data)
	_ = unix.Munmap(data)
}
