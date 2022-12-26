package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ServerConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"loglevel"`
}

type StoreConfig struct {
	Host    string `yaml:"host"`
	DBName  string `yaml:"dbname"`
	SSLMode string `yaml:"sslmode"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
	StoreConfig  StoreConfig  `yaml:"StoreConfig"`
}

func New() *Config {
	return &Config{
		ServerConfig: ServerConfig{"", "", ""},
		StoreConfig:  StoreConfig{"", "", ""},
	}
}

func (c *Config) TakeFromFile(path string) (*Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
