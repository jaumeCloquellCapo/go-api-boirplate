package migrations

import (
	"ApiRest/provider"
	"flag"
	"github.com/pressly/goose"
	"log"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations/", "directory with migration files")
)

// HandlerMigrations ...
func HandlerMigrations(args []string, db *provider.DbStore) {

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Run(args[0], db.DB, "./migrations/"); err != nil {
		log.Fatalf("Error: Migration %v: %v", args[0], err)
	}

}
