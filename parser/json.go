package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func YAMLtoJSON() {
	filePath := "test.yml"

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var yamlMap map[string]interface{}
	err = yaml.Unmarshal([]byte(yamlFile), &yamlMap)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	jsonData, err := json.MarshalIndent(yamlMap, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}