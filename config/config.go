package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}

	ServerConfig struct {
		Port  string
		Debug bool
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		DBName   string
		UserName string
		Password string
	}
)

func LoadConfig(filename string) (Config, error) {
	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config/%s", filename))
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}

func LoadConfigPath(path string) (Config, error) {
	// Load Config
	v := viper.New()

	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, errors.New("config file not found")
		}
		return Config{}, err
	}

	// Parse Config
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return Config{}, err
	}

	return c, nil
}
