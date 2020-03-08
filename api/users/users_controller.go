package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/db"
)

//CreateUser adds a new user
func CreateUser(c *gin.Context) {

	user := User{}
	c.BindJSON(&user)

	text := `INSERT INTO users
	(username, email, password)
	VALUES($1, $2, $3)
	RETURNING user_id
	, username
	, email
	, status_id`

	err := db.DB.QueryRow(text, user.Username, user.Email, user.Password).Scan(&user.UserID, &user.Username, &user.Email, &user.StatusID)
	utils.HandleError(err)

	//referencing the user returns the values (&user)
	//*user doesn't work though, not sure why, should be the other
	//way around
	c.JSON(200, gin.H{"data": user})
}

//GetUser will return a specific user
func GetUser(c *gin.Context) {
	name := c.Param("name")
	//to send a string use Sprintf, the f designator is needed for any formatted strings
	message := fmt.Sprintf("hello %s", name)
	c.JSON(200, gin.H{
		"message": message,
	})
}
