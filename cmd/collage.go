package cmd

import (
	"github.com/lcmps/hippyfm/app"
	"github.com/spf13/cobra"
)

var collage = &cobra.Command{
	Use: "collage",
	RunE: func(cmd *cobra.Command, args []string) error {

		config, err := app.InitConfig()
		if err != nil {
			return err
		}

		conn := app.InstanceAPI(config.Key, config.Secret)

		app.GetAlbumsByPeriod(conn, "luka1498", "7day", 9)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(collage)
}
