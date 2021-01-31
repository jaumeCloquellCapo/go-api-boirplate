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

// Use config file from the flag.
var cfgFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "file", "dev.env", "Config file")
}

func initConfig() {
	// Use config file from the flag.

	if cfgFile != "" {
		if err := godotenv.Load(cfgFile); err != nil {
			log.Fatalf("Error loading %v", cfgFile)
		}
	} else {
		if err := godotenv.Load("dev.env"); err != nil {
			log.Fatalf("Error loading %v", cfgFile)
		}
	}
}
