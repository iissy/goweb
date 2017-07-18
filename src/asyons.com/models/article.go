package models

import "html/template"

// ArticleListResult is hello
type ArticleListResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Article
}

// Article is hello
type Article struct {
	ID          string
	Subject     string
	Keyword     string
	Description string
	Body        template.HTML
	Picture     string
	USID        string
	NickName    string
	Status      int32
	AddDate     string
	Visited     string
}
