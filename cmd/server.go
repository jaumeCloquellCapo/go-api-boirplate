package cmd

import (
	"ApiRest/dic"
	"ApiRest/route"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			dic.InitContainer()
			router := route.Setup()
			router.Run(":" + os.Getenv("APP_PORT"))
		},
	})
}
