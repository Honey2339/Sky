package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func YAMLtoJSON(filePath string) (RootConfigSchema, error) {
	var config RootConfigSchema

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	return config, err
}