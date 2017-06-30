package github

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	var result IssuesSearchResult
	result.TotalCount = 25
	result.Items = []*Issue{}
	for i := 0; i < result.TotalCount; i++ {
		var issue Issue
		issue.Title = "i am title"
		issue.State = "true"
		issue.HTMLURL = "https://asyons.com"
		issue.Number = i + 1
		issue.User = &User{"jimmy", "https://asyons.com"}
		result.Items = append(result.Items, &issue)
	}

	return &result, nil
}
