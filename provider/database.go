package provider

import (
	"ApiRest/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB(config config.Config) (db *sql.DB, err error) {
	dataSourceName := fmt.Sprintf(config.Database.Username + ":" + config.Database.Password + "@/" + config.Database.Database)
	db, err = sql.Open("mysql", dataSourceName)
	return
}
