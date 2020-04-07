package auth

//Credentials models a set of login credentials
type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//User describes an application user returned on login
type User struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	StatusID uint   `json:"statusID"`
}
