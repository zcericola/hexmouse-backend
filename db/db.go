package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq" //postgres driver
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/config"
)

//DB is the Postgres DB
var DB *sql.DB

//Cache is the Redis Cache
var Cache redis.Conn

//createDBURL is the formatted connection string pulled from the dbconfig
func createDBURL(dbConfig *config.DBConfig) string {
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

// //createRDURL is the formatted connection string for redis
// func createRDURL(rdConfig *config.RDConfig) string {

// }

//InitDBConn connects to the database
func InitDBConn() {
	DBConfig := config.GenerateDBConfig()
	DBConnectionString := createDBURL(DBConfig)
	var err error
	instance, err := sql.Open("postgres", DBConnectionString)
	if err = instance.Ping(); err != nil {
		log.Panic(err)
	}
	DB = instance
	log.Print("Db successfully connected.")
}

//Close will close the database connection
func Close() {
	defer DB.Close()
	defer Cache.Close()
	log.Print("Connections closed after main.")

}

//Query will interact with the database
func Query() {

}

//createRedisPool will make a pool for redis connections
func createRedisPool() *redis.Pool {
	redisURL, _ := os.LookupEnv("REDIS_URL")
	// rdPort, _ := os.LookupEnv("RDPORT")
	return &redis.Pool{
		//max # of idle connections
		MaxIdle: 80,
		//max number of active connections at one time
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	pong, err := c.Do("PING")
	utils.HandleError(err)

	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	_, err = redis.String(pong, err)
	utils.HandleError(err)
	return err
}

//InitRedisCache will start redis (used for sessions store)
func InitRedisCache() {
	//Initialize the redis connection to a local redis instance
	pool := createRedisPool()
	conn := pool.Get()
	err := ping(conn)
	utils.HandleError(err)
	Cache = conn

	log.Print("Redis Cache successfully connected.")
}
