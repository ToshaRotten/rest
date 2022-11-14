package apiserver

import (
	"io/ioutil"

	"github.com/ToshaRotten/rest/store"
	"gopkg.in/yaml.v3"
)

// Config default struct
type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"loglevel"`
	Store    *store.Config
}

//New config - create a new config with parameters: { host:"localhost" port:":8080" }
func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

// Take config from file
func TakeConfigFile(path string) (*Config, error) {
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
