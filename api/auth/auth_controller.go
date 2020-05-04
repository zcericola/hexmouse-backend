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

	GenerateSession(user.Username, c)

	c.JSON(200, gin.H{
		"message": "User successfully logged in.",
		"data":    user,
	})
	return
}

//LogoutUserHandler allows a user to logout and destroys their active session
func LogoutUserHandler(c *gin.Context) {
	//get user information from the session cookie
	sessionToken, err := c.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User unathorized.",
			})
			return
		}
	}

	err = DestroySession(sessionToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to destroy session.",
		})
		return
	}

	/*requests that the browser sets a new cookie header with a maxAge of zero and no value
	to ensure that the user does not have a stale cookie*/
	SetCookieForUser("", 0, c)

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully logged out.",
	})
	return
	// c.Redirect(http.StatusFound, "http://localhost:3002/auth/login")
	// return

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
			"message": "Request failed.",
		})
		return
	}

	//go to the Redis Cache and retrieve token
	username, err := redis.String(db.Cache.Do("GET", oldSessionToken))

	//create a new session for user
	newSessionToken, err := RefreshSession(username, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create session.",
		})
	}

	//destroy the older session
	err = DestroySession(oldSessionToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to destroy session.",
		})
	}

	SetCookieForUser(newSessionToken, 120, c)

	c.JSON(http.StatusOK, gin.H{
		"message": "Session successfully refreshed.",
	})
	return
}
