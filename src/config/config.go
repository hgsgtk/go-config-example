package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// DBConfig represents database connection configuration information.
type DBConfig struct {
	User     string `toml:"user"`
	Password string
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
}

// Config represents application configuration.
type Config struct {
	DB DBConfig `toml:"database"`
}

// NewConfig return configuration struct.
func NewConfig(path string, appMode string) (Config, error) {
	var conf Config

	confPath := path + appMode + ".toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return conf, err
	}

	conf.DB.Password = os.Getenv("DB_PASSWORD")

	return conf, nil
}
