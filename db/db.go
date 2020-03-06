package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //postgres driver
	"github.com/zcericola/hexmouse-backend/config"
)

//DB is an instance of the database
type DB *DB

//URL is the formatted connection string pulled from the dbconfig
func URL(dbConfig *config.DBConfig) string {
	config.GenerateDBConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?sslmode=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}

//Init connects to the database
func Init() *sql.DB {
	DBConfig := config.GenerateDBConfig()
	DBConnectionString := URL(DBConfig)
	db, err := sql.Open("postgres", DBConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
