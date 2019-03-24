package models

// User is hello
type User struct {
	ID            int
	UserID        string
	UserName      string
	Password      string
	RegDate       string
	LastLoginDate string
	Status        int
}

// Users is struct
type Users struct {
	PageArgs PageArgs
	Items    []*User
}
