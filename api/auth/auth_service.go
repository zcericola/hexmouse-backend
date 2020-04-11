package auth

import (
	"fmt"
	"log"

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

//GenerateSession will create or update a session for a user
func GenerateSession(username string) string {
	//creates a random session token
	sessionToken := uuid.NewV4().String()
	fmt.Print("sessionToken gen: ", sessionToken)
	_, err := db.Cache.Do("SETEX", sessionToken, "120", username)
	utils.HandleError(err)

	//Todo: Set Cookie here

	return username
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

	//Generate or renew session for user
	GenerateSession(user.Username)

	return User{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Password: "",
		StatusID: user.StatusID,
	}
}
