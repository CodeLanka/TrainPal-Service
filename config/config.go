package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var configMap map[string]string

func InitConfigs(configPath string) {
	var err error
	configMap, err = godotenv.Read(configPath)
	if err != nil {
		log.Fatal("Error reading configs ", err)
	}
	log.Info("Configs load success")
}

func GetConfig(configName string) string {
	value := configMap[configName]
	if value == "" {
		log.Fatal("Invalid Config Read ", configName)
	}
	return value
}
