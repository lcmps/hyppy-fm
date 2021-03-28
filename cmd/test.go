package cmd

import (
	"fmt"

	"github.com/lcmps/hippyfm/app"
	"github.com/spf13/cobra"
)

var test = &cobra.Command{
	Use: "test",
	RunE: func(cmd *cobra.Command, args []string) error {

		config, err := app.InitConfig()
		if err != nil {
			return err
		}

		conn := app.InstanceAPI(config.Key, config.Secret)

		a, _ := app.GetTagsByArtist(conn, "a506f761-2c22-4b2f-8a94-bd748c2c8f75")
		a = app.EnrichTagData(a)
		fmt.Println(a)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(test)
}
