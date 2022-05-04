package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config defines all necessary parameters.
type Config struct {
	Node NodeConfig `mapstructure:"node"`
	DB   DBConfig   `mapstructure:"database"`
	Log  LogConfig  `mapstructure:"log"`
}

// NodeConfig wraps all node endpoints that are used in this project.
type NodeConfig struct {
	AddressPrefix string `mapstructure:"addr_prefix"`
	APIEndpoint   string `mapstructure:"api"`
}

// DBConfig wraps all required parameters for database connection.
type DBConfig struct {
	Mode string `mapstructure:"mode"`
	DSN  string `mapstructure:"dsn"`
}

// LogConfig wraps all required parameters for database connection.
type LogConfig struct {
	Level string `mapstructure:"level"`
	Dir   string `mapstructure:"dir"`
}

// ParseConfig attempts to read and parse config.yaml from the given path
// An error reading or parsing the config results in a panic.
func ParseConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	var config Config
	viper.Unmarshal(&config)
	return &config
}
