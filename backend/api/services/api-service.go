package services

import (
	"github.com/SwanHtetAungPhyo/api/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Configuration() *models.GatewayConfig {
	var config struct {
		Gateway models.GatewayConfig `yaml:"gateway"`
	}
	file, err := os.Open("api-gateway.yaml")
	if err != nil {
		log.Fatal(err.Error())
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
