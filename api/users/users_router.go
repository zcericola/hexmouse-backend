package users

import (
	"github.com/gin-gonic/gin"
)

//Routes holds all user routes
func Routes(route *gin.Engine) {
	router := route.Group("/users")
	router.POST("/", CreateUser)
	router.GET("/:name", GetUser)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DeleteUser)

}
