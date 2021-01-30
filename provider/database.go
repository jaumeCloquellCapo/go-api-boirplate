package provider

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DbStore struct {
	*sql.DB
}

func InitializeDB() *DbStore {
	//dataSourceName := fmt.Sprintf(core.Database.Username + ":" + core.Database.Password + "@/" + core.Database.Database)
	cnf := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	var err error
	var db *sql.DB

	if db, err = sql.Open("mysql", cnf); err != nil {
		log.Fatal(err)
	}
	if errPing := db.Ping(); errPing != nil {
		log.Fatal(errPing)
	}
	return &DbStore{
		db,
	}
}
