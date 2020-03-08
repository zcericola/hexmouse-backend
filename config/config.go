package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//DBConfig used to connect to DB
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	URI      string
}

//Init will load environment variables
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		log.Print("No .env file found.")
	} else {
		log.Print("Environment vars loaded successfully.")
	}
}

//GenerateDBConfig generates a connection string
func GenerateDBConfig() *DBConfig {
	host, _ := os.LookupEnv("PGHOST")
	port, _ := os.LookupEnv("PGPORT")
	user, _ := os.LookupEnv("PGUSER")
	name, _ := os.LookupEnv("PGDATABASE")
	password, _ := os.LookupEnv("PGPASSWORD")
	URI, _ := os.LookupEnv("PGURI")

	DBConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Name:     name,
		Password: password,
		URI:      URI,
	}

	return &DBConfig
}
