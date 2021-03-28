package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a default configuration struc for this app
type Config struct {
	//App Version
	Version string
	Key     string
	Secret  string
	Port    string
}

// InitConfig initiates this application default configuration
func InitConfig() (*Config, error) {

	config := &Config{
		Version: fmt.Sprintf("%v", viper.Get("VERSION")),
		Key:     fmt.Sprintf("%v", viper.Get("KEY")),
		Secret:  fmt.Sprintf("%v", viper.Get("SECRET")),
		Port:    fmt.Sprintf("%v", viper.Get("PORT")),
	}

	if len(config.Version) == 0 {
		return nil, fmt.Errorf("version must be set")
	}

	if len(config.Key) == 0 {
		return nil, fmt.Errorf("key must be set")
	}

	if len(config.Secret) == 0 {
		return nil, fmt.Errorf("secret must be set")
	}

	return config, nil
}
