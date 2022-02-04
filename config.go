package main

import (
	"os"
	"path"
)

const (
	ConfigDirName = ".taks/"      // The name of the directory for the config file
	ConfigFile    = "config.yaml" // The filename for storing app config data
	DataFile      = "data.db"     // The filename for storing app data
)

// AppConfig stores the user's configuration
// for the application.
type AppConfig struct {
	Path string // Location of the taks data directory
}

func NewAppConfig(p string) (*AppConfig, error) {
	if p == "" {
		var err error
		p, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
	}
	d := path.Join(p, ConfigDirName)
	cfg := &AppConfig{Path: d}
	return cfg, nil
}
