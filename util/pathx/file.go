package pathx

import (
	"os"
	"path/filepath"
)

const (
	NL              = "\n"
	goctlDir        = ".che"
	gitDir          = ".git"
	autoCompleteDir = ".auto_complete"
	cacheDir        = "cache"
)

var goctlHome string

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// GetGoctlHome returns the path value of the goctl, the default path is ~/.goctl, if the path has
// been set by calling the RegisterGoctlHome method, the user-defined path refers to.
func GetGoctlHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			os.Rename(home, home+".old")
			MkdirIfNotExist(home)
		}
	}()
	if len(goctlHome) != 0 {
		home = goctlHome
		return
	}
	home, err = GetDefaultGoctlHome()
	return
}

// GetDefaultGoctlHome returns the path value of the goctl home where Join $HOME with .goctl.
func GetDefaultGoctlHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, goctlDir), nil
}

// GetGitHome returns the git home of goctl.
func GetGitHome() (string, error) {
	goctlH, err := GetGoctlHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(goctlH, gitDir), nil
}
