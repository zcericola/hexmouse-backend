package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //postgres driver
	"github.com/zcericola/hexmouse-backend/config"
)

//DB is the database type
var DB *sql.DB

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
func Init() {
	DBConfig := config.GenerateDBConfig()
	DBConnectionString := URL(DBConfig)
	var err error
	instance, err := sql.Open("postgres", DBConnectionString)
	if err = instance.Ping(); err != nil {
		log.Panic(err)
	}
	DB = instance
	// defer DB.Close()
	log.Print("Db successfully connected.")
}

//Close will close the database connection
func Close() {
	DB.Close()

}

//Query will interact with the database
func Query() {

}
