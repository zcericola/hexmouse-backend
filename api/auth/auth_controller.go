package auth

import (
	"github.com/gin-gonic/gin"
)

//LoginUserHandler allows a user to login to the application
func LoginUserHandler(c *gin.Context) {
	params := Credentials{}
	c.BindJSON(&params)
	user := LoginUser(params)
	c.JSON(200, gin.H{
		"message": "User successfully logged in.",
		"data":    user,
	})
}
