package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq" //postgres driver
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/config"
)

//DB is the Postgres DB
var DB *sql.DB

//Cache is the Redis Cache
var Cache redis.Conn

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
	log.Print("Db successfully connected.")
}

//InitRedisCache will start redis (used for sessions store)
func InitRedisCache() {
	//Initialize the redis connection to a local redis instance
	conn, err := redis.DialURL("redis://localhost")
	utils.HandleError(err)
	Cache = conn
	log.Print(Cache)
	log.Print("Redis Cache successfully connected.")
}

//Close will close the database connection
func Close() {
	DB.Close()

}

//Query will interact with the database
func Query() {

}
