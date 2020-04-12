package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/zcericola/hexmouse-backend/api/utils"
	"github.com/zcericola/hexmouse-backend/db"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword generates a hash from a password string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(bytes), err
}

//CheckPasswordHash will compare the hash and check validity against a password string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}
	return false
}

//CheckForValidUserStatus helps authorize user
func CheckForValidUserStatus(statusID uint) bool {
	//Todo: expand this later
	if statusID == 1 {
		return true
	}
	return false
}

//SetCookieForUser will set the session cookie
func SetCookieForUser(tokenValue string, c *gin.Context) {
	c.SetCookie("session_token", tokenValue, 120, "/", "localhost", false, true)
}

//DeleteSession will remove a session from the cache
func DeleteSession(sessionKey string) (err error) {
	_, err = db.Cache.Do("DEL", sessionKey)
	return err
}

//RefreshSession refreshes a session for an existing user
func RefreshSession(username string, c *gin.Context) (string, error) {
	newSessionToken := uuid.NewV4().String()
	_, err := db.Cache.Do("SETEX", newSessionToken, "120", username)
	return newSessionToken, err
}

//GenerateSession will create an initial session for a user
func GenerateSession(username string, c *gin.Context) {
	//creates a random session token
	sessionToken := uuid.NewV4().String()
	//sets the token in the redis cache with 120 second ttl
	_, err := db.Cache.Do("SETEX", sessionToken, "120", username)
	utils.HandleError(err)

	//check if there is a session_token already
	_, err = c.Cookie("session_token")
	//if no cookie, set the token
	if err != nil {
		SetCookieForUser(sessionToken, c)
	}
}

//LoginUser allows a user to login to the application
func LoginUser(params Credentials) User {
	user := User{}
	selectStatement := `
		SELECT u.username
		, u.email
		, u.Password
		, u.status_id
		FROM users u
		WHERE u.username = $1;`

	stmt, err := db.DB.Prepare(selectStatement)
	utils.HandleError(err)

	err = stmt.QueryRow(
		params.Username,
	).Scan(
		&user.Username,
		&user.Email,
		&user.Password,
		&user.StatusID,
	)
	utils.HandleError(err)

	//hash user provided password and compare to db hash
	var isValidPassword bool = CheckPasswordHash(params.Password, user.Password)
	//Clearing hashes to prevent logging
	params.Password = ""
	user.Password = ""
	if isValidPassword == false {
		log.Panic("Invalid Password.")
	}

	//Check if user is active
	var isValidStatus bool = CheckForValidUserStatus(user.StatusID)

	if isValidStatus == false {
		log.Panic("User has been deactivated.")
	}

	return User{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Password: "",
		StatusID: user.StatusID,
	}
}
