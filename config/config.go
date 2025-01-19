package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func ParseConfig(configPath string) Config {
	config := Config{}

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config: %v", err.Error()))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing config: %v", err.Error()))
	}

	return config
}
