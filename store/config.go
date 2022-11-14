package store

// Config ...
type Config struct {
	DatabaseURL string `yaml:"database_url"`
}

// New ...
func NewConfig() *Config {
	return &Config{}
}
