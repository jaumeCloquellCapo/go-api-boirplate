package cmd

import (
	"ApiRest/migrations"
	"ApiRest/provider"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "migration",
		Short: "Run migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migrations.HandlerMigrations(args, provider.InitializeDB())
		},
	})
}
