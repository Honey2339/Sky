package config

import (
	"encoding/json"
	"io"
	"os"
)

func Get_config_data() (RootConfigSchema, error) {
	path, err := Get_config()
	var data RootConfigSchema

	if err != nil {
		return data, err
	}

	file, err := os.Open(path)

	if err != nil {
		return data, err
	}

	defer file.Close()

	byteValue, _ := io.ReadAll(file)


	json.Unmarshal(byteValue, &data)

	return data, nil
}