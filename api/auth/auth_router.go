package auth

import (
	"github.com/gin-gonic/gin"
)

//Routes holds all auth routes
func Routes(route *gin.Engine) {
	router := route.Group("/auth")
	router.POST("/", LoginUserHandler)
	router.GET("/refresh", RefreshSessionHandler)
}
