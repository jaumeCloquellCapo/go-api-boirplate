package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
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

	switch env {
	case "dev":
		if err := godotenv.Load(fmt.Sprintf("%v.env", env)); err != nil {
			log.Fatalf("Error loading %v.env", env)
		}
	case "pro":
		if err := godotenv.Load(fmt.Sprintf("%v.env", env)); err != nil {
			log.Fatalf("Error loading %v.env", env)
		}
	default:
		if err := godotenv.Load(fmt.Sprintf("%v.env", env)); err != nil {
			log.Fatalf("Error loading %v.env", env)
		}
	}
}
