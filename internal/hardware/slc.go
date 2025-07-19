//go:build darwin && arm64

package hardware

import (
	"os/exec"
	"strconv"
	"strings"
)

// DetectSLCMiB returns the size of the system level cache in MiB on
// Apple silicon machines. It returns 0 if detection fails.
func DetectSLCMiB() int {
	out, err := exec.Command("sysctl", "-n", "hw.l3cachesize").Output()
	if err != nil {
		return 0
	}
	n, _ := strconv.Atoi(strings.TrimSpace(string(out)))
	return n / (1024 * 1024)
}
