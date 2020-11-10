package cmd

import (
	"ApiRest/dic"
	"ApiRest/route"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	var serverPort string
	defaultServerPort := os.Getenv("APP_PORT")
	serverCmd.PersistentFlags().StringVar(&serverPort, "port", defaultServerPort, "App port")

	//cobra.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		dic.InitContainer()

		router := route.Setup()
		router.Run(":" + os.Getenv("APP_PORT"))
	},
}
