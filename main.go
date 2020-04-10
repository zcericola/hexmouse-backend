package main

import (
	"log"

	"github.com/zcericola/hexmouse-backend/config"
	"github.com/zcericola/hexmouse-backend/db"
	"github.com/zcericola/hexmouse-backend/server"
)

func init() {
	config.Init()
	db.Init()
	db.InitRedisCache()
	log.Print("All init actions completed.")

}

func main() {
	server.Init()
	defer db.Close()
}
