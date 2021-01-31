package cmd

import (
	"ApiRest/dic"
	"ApiRest/route"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			environment, _ := cmd.Flags().GetString("env")
			switch environment {
			case "dev":
				if err := godotenv.Load(fmt.Sprintf("%v.env", environment)); err != nil {
					log.Fatalf("Error loading %v.env", environment)
				}
			case "pro":
				if err := godotenv.Load(fmt.Sprintf("%v.env", environment)); err != nil {
					log.Fatalf("Error loading %v.env", environment)
				}
			default:
				if err := godotenv.Load(fmt.Sprintf("%v.env", environment)); err != nil {
					log.Fatalf("Error loading %v.env", environment)
				}
			}

			container := dic.InitContainer()
			router := route.Setup(container)
			router.Run(":" + os.Getenv("APP_PORT"))
		},
	})
}
