package cmd

import (
	"github.com/lcmps/hippyfm/app"
	"github.com/spf13/cobra"
)

var web = &cobra.Command{
	Use: "web",
	RunE: func(cmd *cobra.Command, args []string) error {

		config, err := app.InitConfig()
		if err != nil {
			return err
		}
		conn := app.InstanceAPI(config.Key, config.Secret)

		app.Host(conn)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(web)
}
