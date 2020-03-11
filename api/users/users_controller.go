package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/db"
)

//CreateUser adds a new user
func CreateUser(c *gin.Context) {

	params := User{}
	c.BindJSON(&params)

	newUser := User{}

	insertStatement := `INSERT INTO users
	(username, email, password)
	VALUES($1, $2, $3)
	RETURNING user_id
	, username
	, email
	, status_id`

	stmt, err := db.DB.Prepare(insertStatement)

	err = stmt.QueryRow(
		params.Username,
		params.Email,
		params.Password,
	).Scan(&newUser.UserID,
		&newUser.Username,
		&newUser.Email,
		&newUser.StatusID,
	)
	utils.HandleError(err)

	c.JSON(200, gin.H{"data": newUser})
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
