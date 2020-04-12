package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/zcericola/hexmouse-backend/db"
)

//LoginUserHandler allows a user to login to the application
func LoginUserHandler(c *gin.Context) {
	params := Credentials{}
	c.BindJSON(&params)
	user := LoginUser(params)

	//Generate or renew session for user
	GenerateSession(user.Username, c)

	c.JSON(200, gin.H{
		"message": "User successfully logged in.",
		"data":    user,
	})
}

//RefreshSessionHandler will renew a logged in user's session token
func RefreshSessionHandler(c *gin.Context) {
	oldSessionToken, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User unauthorized.",
			})
			return
		}
		c.JSON(400, gin.H{
			"message": "Bad request.",
		})
		return
	}

	//go to the Redis Cache and retrieve token
	username, err := redis.String(db.Cache.Do("GET", oldSessionToken))

	//create a new session for user
	newSessionToken, err := RefreshSession(username, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error.",
		})
	}

	//delete the older session
	err = DeleteSession(oldSessionToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error.",
		})
	}

	SetCookieForUser(newSessionToken, c)

	c.JSON(http.StatusOK, gin.H{
		"message": "Session successfully refreshed.",
	})
}
