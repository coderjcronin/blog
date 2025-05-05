package config

import (
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	ex, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	returnPath := filepath.Join(ex, ".gatorconfig.json")

	return returnPath, nil
}
