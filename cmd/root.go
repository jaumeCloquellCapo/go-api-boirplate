package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "api",
	Long:  `CRUD API Voicemod`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var env string

func Execute() {
	//rootCmd.Use = viper.GetString("APP_COMMAND")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&env, "env", "dev", "Environment name.")
}

func initConfig() {
	env, _ := rootCmd.Flags().GetString("env")
	if env != "" {
		fmt.Println("env:", env)
	}
}
