package github

// IssuesSearchResult is hello
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Article
}

// User is hello
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// Article is hello
type Article struct {
	ID          string
	Subject     string
	Keyword     string
	Description string
	Body        string
	Picture     string
	USID        string
	NickName    string
	Status      int32
	AddDate     string
	Visited     string
}
