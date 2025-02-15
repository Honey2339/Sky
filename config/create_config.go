package config

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
)

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Create_config() error {
	configPath, err := Get_config()
	if err != nil {
		return err
	}

	// You need a add_config function here
	// So that when a config is created, You can ask user to fill the data
	// Add_config()

	config := Config{
		Key:   "exampleKey",
		Value: "exampleValue",
	}

	file, err := os.Create(configPath)
	
	if err != nil {
		return err
	}
	
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return err
	}

	log.Info("Config file created at:", configPath)

	return nil
}