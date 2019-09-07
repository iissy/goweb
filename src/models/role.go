package models

// Role is yes
type Role struct {
	ID          int
	RoleName    string
	Status      int
	CreateTime  string
	UpdatedTime string
}

// Functionality is yes
type Functionality struct {
	ID          int
	Funname     string
	FunType     string
	Controller  string
	CreateTime  string
	UpdatedTime string
}

// RoleFunctionMapping is yes
type RoleFunctionMapping struct {
	FunID  int
	RoleID int
	Toggle bool
}

// Roles is struct
type Roles struct {
	PageArgs PageArgs
	Items    []*Role
}

// Functionalities is struct
type Functionalities struct {
	PageArgs PageArgs
	Items    []*Functionality
}
