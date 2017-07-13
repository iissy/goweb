package github

import (
	"database/sql"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	var result IssuesSearchResult
	result.TotalCount = 25
	result.Items = []*Issue{}

	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.188:3306)/iissy?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select LastName,FirstName,Address from Persons limit ?", 10)
	checkErr(err)
	var i = 0
	for rows.Next() {
		var issue Issue
		err = rows.Scan(&issue.Title, &issue.HTMLURL, &issue.State)
		checkErr(err)
		issue.State = "true"
		issue.Number = i + 1
		issue.User = &User{issue.Title, issue.HTMLURL}
		result.Items = append(result.Items, &issue)
		i++
	}

	return &result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
