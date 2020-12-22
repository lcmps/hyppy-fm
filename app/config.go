package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a default configuration struc for this app
type Config struct {
	//App Version
	Version string
}

// InitConfig initiates this application default configuration
func InitConfig() (*Config, error) {

	config := &Config{
		Version: viper.GetString("version"),
	}

	if len(config.Version) == 0 {
		return nil, fmt.Errorf("CRM Endpoint must be set")
	}

	return config, nil
}
