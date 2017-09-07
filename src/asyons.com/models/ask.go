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

// AskSimpleListView is list for embeded
type AskSimpleListView struct {
	ID      string
	Subject string
}

// AskView is for view model
type AskView struct {
	Subject string
	Body    template.HTML
	AddDate string
	List    []*AskSimpleListView
}
