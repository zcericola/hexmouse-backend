package auth

import (
	"github.com/gin-gonic/gin"
)

//Routes holds all auth routes
func Routes(route *gin.Engine) {
	router := route.Group("/auth")
	router.POST("/login", LoginUserHandler)
	router.POST("/logout", LogoutUserHandler)
	router.GET("/refresh", RefreshSessionHandler)
}
