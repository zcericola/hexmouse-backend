package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zcericola/hexmouse-backend/api/auth"
	"github.com/zcericola/hexmouse-backend/api/snippets"
	"github.com/zcericola/hexmouse-backend/api/users"
)

//PORT defaults to localhost:3002
const PORT string = "localhost:3002"

//Init will start the server
func Init() {
	router := gin.Default()
	auth.Routes(router)
	users.Routes(router)
	snippets.Routes(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(PORT) //listen and serve on port 3002
}
