package models

import "html/template"

// Comment is yes
type Comment struct {
	ID       int
	UserID   int
	NickName string
	AskID    int
	Body     template.HTML
	AddDate  string
}

// Comments is hello
type Comments struct {
	TotalCount int `json:"total_count"`
	Items      []*Comment
}
