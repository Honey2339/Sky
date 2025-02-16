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
	
	var config RootConfigSchema
	var file *os.File
	var checkFile bool = false
	
	if err != nil {
		return err
	}
	
	_, err = os.Stat(configPath)
	
	if err == nil {
		checkFile = true
	}
	
	if(!checkFile){
		file, err = os.Create(configPath)

		if err != nil {
			return err
		}
	
		config , err = Add_config()
		
		if err != nil {
			return err
		}
	} else {
		log.Error("Config File Already Exist")
	}
	
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return err
	}

	log.Print("Config file created at:", configPath)

	return nil
}