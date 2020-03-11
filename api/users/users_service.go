package users

//SanitizeUser cleans up user data
func SanitizeUser(u User) User {
	return User{u.UserID, u.Username, u.Email, "", u.StatusID}
}
