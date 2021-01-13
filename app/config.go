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
	Env     string
}

// InitConfig initiates this application default configuration
func InitConfig() (*Config, error) {

	config := &Config{
		Version: viper.GetString("version"),
		Key:     viper.GetString("key"),
		Secret:  viper.GetString("secret"),
		Env:     viper.GetString("env"),
	}

	if len(config.Version) == 0 {
		return nil, fmt.Errorf("version must be set")
	}

	if len(config.Key) == 0 {
		return nil, fmt.Errorf("Key must be set")
	}

	if len(config.Secret) == 0 {
		return nil, fmt.Errorf("Secret must be set")
	}

	if len(config.Env) == 0 {
		return nil, fmt.Errorf("Env must be set")
	}

	return config, nil
}
