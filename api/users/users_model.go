package users

//User describes an application user
type User struct {
	UserID   uint
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` //don't encode, don't return
	StatusID int
}
