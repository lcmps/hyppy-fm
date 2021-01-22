package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var cfgPath = "./configs"

// Execute the application.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

var rootCmd = &cobra.Command{
	Short: "hippy-fm backend",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("version"))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./configs/default.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("default")
		viper.AddConfigPath(cfgPath)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to read config from file")
		// logrus.WithError(err).Fatal("Unable to read config from file")
	}
}
