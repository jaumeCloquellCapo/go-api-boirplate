package provider

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func InitializeDB() (db *sql.DB) {
	//dataSourceName := fmt.Sprintf(config.Database.Username + ":" + config.Database.Password + "@/" + config.Database.Database)
	cnf := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	var err error

	if db, err = sql.Open("mysql", cnf); err != nil {
		log.Fatal(err)
	}
	if errPing := db.Ping(); errPing != nil {
		log.Fatal(errPing)
	}
	return
}
