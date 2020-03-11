package users

//User describes an application user
type User struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	StatusID int    `json:"statusID"`
}
