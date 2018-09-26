package models

// Course is hello
type Course struct {
	ArticleItems []*Article
}

// Picture is struct
type Picture struct {
	ID      string
	URL     string
	Subject string
}

// Pictures is struct
type Pictures struct {
	TotalCount int
	Items      []*Picture
}
