package users

import (
	"fmt"
	"io/ioutil"

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
	utils.HandleError(err)

	err = stmt.QueryRow(
		params.Username,
		params.Email,
		params.Password,
	).Scan(&newUser.UserID,
		&newUser.Username,
		&newUser.Email,
		&newUser.StatusID,
	)

	c.JSON(200, gin.H{
		"message": "User successfully created.",
		"data":    newUser,
	})
}

//GetUser will return a specific user
func GetUser(c *gin.Context) {
	name := c.Param("name")
	//to send a string use Sprintf, the f designator is needed for any formatted strings
	message := fmt.Sprintf("hello %s", name)
	//TODO: FIX THIS TO MIRROR RETURN FORMAT OF OTHER ENDPOINTS
	c.JSON(200, gin.H{
		"message": "User successfully retrieved.",
		"data":    message,
	})
}

//UpdateUser will update some user attributes
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id") //proper way to get URL param
	params := User{}
	c.BindJSON(&params)

	//next two lines show how to log from body
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("username", string(body)) //must convert buff to string

	updatedUser := User{}

	updateStatement := `
		UPDATE users
		SET username = $2
		WHERE user_id = $1
		RETURNING
		user_id
		, username
		, email`

	stmt, err := db.DB.Prepare(updateStatement)
	utils.HandleError(err)

	err = stmt.QueryRow(
		idParam,
		params.Username,
	).Scan(
		&updatedUser.UserID,
		&updatedUser.Username,
		&updatedUser.Email,
	)

	c.JSON(200, gin.H{
		"message": "User successfully updated.",
		"data":    updatedUser,
	})
}

//DeleteUser will delete a user profile
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deleteStatement := `
		DELETE FROM users
		WHERE user_id = $1
	`
	stmt, err := db.DB.Prepare(deleteStatement)
	utils.HandleError(err)

	_, err = stmt.Exec(id)
	utils.HandleError(err)

	c.JSON(200, gin.H{
		"message": "User successfully deleted.",
		"data":    id,
	})
}
