package models

import (
	"html/template"
)

// Article is struct
type Article struct {
	ID          string
	PostType    string
	Subject     string
	Picture     string
	Description string
	Body        template.HTML
	UserID      int
	NickName    string
	Origin		string
	AddDate     string
	Visited     int
	List        []*Article
}

// Articles is struct
type Articles struct {
	PageArgs PageArgs
	Items    []*Article
}
