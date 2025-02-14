package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

func ConfigData() {
	filePath := "test.yml"
	config, err := YAMLtoJSON(filePath)
	if err != nil {
		log.Fatalf("Error processing YAML: %v", err)
	}
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation: %s\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("Validation passed!")
	}
	jsonData, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(jsonData))
}