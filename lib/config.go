package lib

import (
	"os"
	"path"
)

const (
	ConfigDirName = ".taks/"
)

// GetConfigDir returns the path to the config directory
// where the database should be stored. `p` is the path
// to the parent directory, in which the config directory
// will be named after `ConfigDirName`. If `p` is empty,
// the config directory will be created in the user's
// home directory (per `os.UserHomeDir()`).
func GetConfigDir(p string) (string, error) {
	if p != "" {
		return path.Join(p, ConfigDirName), nil
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homedir, ConfigDirName), nil
}
