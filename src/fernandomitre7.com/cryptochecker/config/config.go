package config

import (
	"encoding/json"
	"os"

	"fernandomitre7.com/cryptochecker/logger"
)

// Configuration is the structure holding projects configuration environment variables
type Configuration struct {
	Port        int    `json:"port"`
	Environment string `json:"environment"`
	Debug       bool   `json:"debug"`
	BitsoAPIURL string `json:"bitso_api_url"`
}

var (
	_conf *Configuration
)

// LoadConfiguration loads project configuration from a file and from specified env vars
func LoadConfiguration(filePath string) (conf *Configuration, err error) {
	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		logger.Error("Error openning config file %s: %v", filePath, err.Error())
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&conf); err != nil {
		logger.Error("Error decoding config file %v", err.Error())
		return nil, err
	}
	_conf = conf
	return

	// TODO: Set into Configuration the values that are stored in env vars
	//	 Values in env vars are for security, passwords and stuff
	//configuration.Connection_String = os.Getenv("Connection_String")
}

// GetConfiguration returns the loaded configuration
func GetConfiguration() (conf *Configuration) {
	return _conf
}
