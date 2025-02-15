package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/charmbracelet/log"
)

func Get_config()(string, error) {
	var configDir string
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "",err
	}

	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(homeDir, "Appdata", "Roaming", "skyrp")
	case "linux", "darwin":
		configDir = filepath.Join(homeDir, ".config", "skyrp")
	default:
		log.Error("Unsupported OS")
		return "", err 
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}