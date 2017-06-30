package github

import "time"

// IssuesSearchResult is hello
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is hello
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User is hello
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
