package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zcericola/hexmouse-backend/api/users"
)

//PORT defaults to localhost:3002
const PORT string = ":3002"

func getUser(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name) //sends it back like res.status(200).json...
}

//Init will start the server
func Init() {
	router := gin.Default()
	router.POST("/users", users.CreateUser)
	router.GET("/users/:name", getUser)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(PORT) //listen and serve on port 3002

}
