package provider

import (
	"ApiRest/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitializeDB(config config.Config) (db *sql.DB, err error) {
	//dataSourceName := fmt.Sprintf(config.Database.Username + ":" + config.Database.Password + "@/" + config.Database.Database)
	cnf := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Database)
	db, err = sql.Open("mysql", cnf)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return
}
