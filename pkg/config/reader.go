package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func ReadConfig[T *ConfigDb | *Server](fieldTitle string, structure T) T {
	config := &structure
	stringMap := viper.GetStringMap(fieldTitle)
	err := mapstructure.Decode(stringMap, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return *config
}
