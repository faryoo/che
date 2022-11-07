package env

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/faryoo/che/vars"
)

func LookPath(xBin string) (string, error) {
	suffix := getExeSuffix()
	if len(suffix) > 0 && !strings.HasSuffix(xBin, suffix) {
		xBin = xBin + suffix
	}

	bin, err := exec.LookPath(xBin)
	if err != nil {
		return "", err
	}
	return bin, nil
}

func getExeSuffix() string {
	if runtime.GOOS == vars.OsWindows {
		return ".exe"
	}
	return ""
}

// CanExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
func CanExec() bool {
	switch runtime.GOOS {
	case vars.OsJs, vars.OsIOS:
		return false
	default:
		return true
	}
}
