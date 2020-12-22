package cmd

import (
	"fmt"

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

		fmt.Println(config.Version)
		app.Host()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(web)
}
