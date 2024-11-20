package models

import (
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

func Configuration() *GatewayConfig {
	var config struct {
		Gateway GatewayConfig `yaml:"gateway"`
	}

	file, err := os.Open("api-gateway.yaml")
	if err != nil {
		log.Fatalf("Config file not found. Please mount the config file to /config/api-gateway.yaml: %s", err.Error())
		return nil
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &config.Gateway
}
