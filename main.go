package main

import (
	"github.com/zcericola/hexmouse-backend/config"
	"github.com/zcericola/hexmouse-backend/db"
	"github.com/zcericola/hexmouse-backend/server"
)

func init() {
	config.Init()

}

func main() {
	server.Init()
	db.Init()
}
