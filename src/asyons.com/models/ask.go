package models

import "html/template"

// AskList is hello
type AskList struct {
	TotalCount int `json:"total_count"`
	Items      []*Ask
}

// Ask is yes
type Ask struct {
	ID          string
	Subject     string
	Description string
	Body        template.HTML
	UserID      int32
	NickName    string
	AddDate     string
	Visited     int32
	Replied     int32
}
